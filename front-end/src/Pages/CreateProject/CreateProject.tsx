import React, { ChangeEvent, useContext, useState} from 'react'
import { AuthContext } from '../../context/AuthContext'
import { ProjectBase } from "../../types/project"

interface FormFieldsError {
  nameErr: string
  descriptionErr: string
  categoryErr: string
  tagsErr: string
}

const CreateProject:React.FC = () => {
  const { loggedIn } = useContext(AuthContext);
  const [showLogIn, setShowLogIn ] = useState<boolean>(false)
  const [showSignUp, setShowSignUp ] = useState<boolean>(false)
  const [formData, setFormData] = useState<ProjectBase>({
    name       : '',
    description: '',
    category   : '',
    tags       : [],
  })
  const [formFieldError, setFormFieldError] = useState<FormFieldsError>({nameErr: "", descriptionErr: "", categoryErr: "", tagsErr: ""});
  const categories = ['business', 'software engineering', 'art'] // TODO get from backend?

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

  const onbSubmit = () => {
    if (!loggedIn) {
      
    }
  }

  return (
    <>
    <div>
      <h2>Create a new project</h2>
    </div>
    <div id='LogInModal-input-container'>
        <input type="text"  name="name" placeholder="name" value={formData.name} onChange={handleChange}/>
        <div className='missingField'>
            {formFieldError.nameErr}
        </div>
    
        <input type="text" name="description" placeholder="description" value={formData.description} onChange={handleChange}/>
        <div className='missingField'>
            {formFieldError.descriptionErr}
        </div>

        <label>select category</label>
        <select onChange={selectCategory}>
          {categories.map((value) => (
              <option key={value} value={value}> {value} </option>
            ))
          }
        </select>
        <div className='missingField'>
            {formFieldError.categoryErr}
        </div>

        <input type="text" name="tags" placeholder="tags" value={formData.tags} onChange={handleChange}/>
        <div className='missingField'>
            {formFieldError.tagsErr}
        </div>
        <button type="submit" onClick={()=>{}}>Create</button>
        <div className='missingField'>
          submit error
        </div>
    </div>
    </>

  )
}

export default CreateProject