import React, { ChangeEvent, useContext, useState} from 'react'
import { AuthContext } from '../../context/AuthContext'
import { ProjectBase } from "../../types/project"
import PromptAccount from '../modals/PromptAccount'
import axiosBase from '../../config/axiosConfig'
import '../modals/Modal.css'
import { useNavigate } from 'react-router-dom'
import { AxiosResponse } from 'axios'

interface FormFieldsError {
  nameErr: string
  descriptionErr: string
  categoryErr: string
  tagsErr: string
}

interface ProjectForm {
  name: string
  description: string
  category: string
  tags: string
}

interface PostData extends ProjectBase {
  ownerId: string
}

const CreateProject:React.FC = () => {
  const navigate = useNavigate()
  const { loggedIn, userID } = useContext(AuthContext)
  const [showPromptAccount, setShowPromptAccount ] = useState<Boolean>(false)

  const [formData, setFormData] = useState<ProjectForm>({
    name       : '',
    description: '',
    category   : '',
    tags       : '',
  })
  const noErrorObj: FormFieldsError = {nameErr: "", descriptionErr: "", categoryErr: "", tagsErr: ""}
  const [formFieldError, setFormFieldError] = useState<FormFieldsError>(noErrorObj);
  const [submitError, setSubmitError] = useState<String>("")

  const categories = ['business', 'software engineering', 'art'] // TODO get from backend?

  const checkFormError = () => {
    let newState: FormFieldsError = { ...noErrorObj }
    const {name, description, tags} = formData
    // TODO define these in some config
    const nameMin=7, descriptionMin=7, tagsArrMin=3
    const nameMax=100, descriptionMax=1000, tagMax=30, tagsArrMax=10

    if (name === '') {
      newState.nameErr = 'Name cannot be empty'
    } else if (name.length < nameMin) {
      newState.nameErr = `Name must be at least ${nameMin} characters`
    } else if (name.length > nameMax) {
      newState.nameErr = `Name cannot exceed ${nameMax} characters`
    }

    if (description === '') {
      newState.descriptionErr = 'description cannot be empty'
    } else if (description.length < descriptionMin) {
      newState.descriptionErr = `description must be at least ${descriptionMin} characters`
    } else if (description.length > descriptionMax) {
      newState.descriptionErr = `description cannot exceed ${descriptionMax} characters`
    }

    const tagArr = tags.split(',').map(e => e.trim())
    if (tagArr.length < tagsArrMin) {
      newState.tagsErr = `must have at least ${tagsArrMin} tags`
    } else if (tagArr.length > tagsArrMax) {
      newState.tagsErr = `can't have more than ${tagsArrMax} tags`
    }
    for (const e of tagArr) {
      if (e.length > tagMax) {
        newState.tagsErr = `a tag can't have more than ${tagMax} characters`
      }
    }
    setFormFieldError(newState)
    let errorHappened = false
    Object.entries(newState).forEach(([k, v]) => {
      if (v !== "") {
        errorHappened = true
      }
    })
  
    return errorHappened
  }

  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData({           
        ...formData,
        [name]: value,
      });
  };

  const selectCategory = (e: ChangeEvent<HTMLSelectElement>) => {
    let newState = formData;
    newState.category = e.target.value;
    setFormData(newState);
  }

  const onSubmit = async () => {

    if (!loggedIn) {
      setShowPromptAccount(true)
      return
    }
    // TODO check form errors
    if (checkFormError()) {
      return
    }
    const tagArr = formData.tags.split(',').map(e => e.trim())

    const postData : PostData = { 
      name:        formData.name, 
      description: formData.description, 
      category:    formData.category, 
      tags:        tagArr, 
      ownerId:     userID
    }
    try {
      const id: AxiosResponse<string> = await axiosBase.post(`/project`, postData);
      navigate(`/project/${id}`, {state: postData})
    } catch (err) {
      setSubmitError("An error occurred")
    }
  }

  return (
    <div>
      <div>
        <h2>Create a new project</h2>
      </div>
      <div id='LogInModal-input-container'>
        <input type="text"  name="name" placeholder="name" value={formData.name} onChange={handleChange}/>
        <div className='errorMessage'>
            {formFieldError.nameErr}
        </div>
    
        <input type="text" name="description" placeholder="description" value={formData.description} onChange={handleChange}/>
        <div className='errorMessage'>
            {formFieldError.descriptionErr}
        </div>

        <label>select category</label>
        <select onChange={selectCategory}>
          {categories.map((value) => (
              <option key={value} value={value}> {value} </option>
            ))
          }
        </select>
        <div className='errorMessage'>
            {formFieldError.categoryErr}
        </div>

        <input type="text" name="tags" placeholder="tags" value={formData.tags} onChange={handleChange}/>
        <div className='errorMessage'>
            {formFieldError.tagsErr}
        </div>
        <button type="submit" onClick={onSubmit  }>Create</button>
        <div className='errorMessage'>
          {submitError}
        </div>
      </div>
      {showPromptAccount &&
        <PromptAccount 
          SetShowPromptAccount={setShowPromptAccount}
          />
      }
    </div>

  )
}

export default CreateProject