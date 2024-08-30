import { ChangeEvent, FC, useState } from "react";
import { User, UserType } from "../../types/user";
import axiosBase from "../../config/axiosConfig";
import { AxiosResponse } from "axios";
import './User.css'

interface UserSettingsForm {
  name        : string
  email       : string
  description : string
  skills      : string
}

interface PostData  {
  name        : string
  email       : string
  description : string
  skills      : string[]
}

interface Props {
  user: UserType
}
const PrivateProfile:FC<Props> = ({ user }) => {
  const [formData, setFormData] = useState<UserSettingsForm>({
    name       : user?.name || '',
    email: (user as User)?.email || '',
    description   : user?.description || '',
    skills       : user?.skills?.join(',') || '',
  })
  
  const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData({           
        ...formData,
        [name]: value,
      })
  }

  const onSubmit = async () => {
    // if (checkFormError(noErrorObj, formData, setFormFieldError)) {
    //   return
    // }
    const tagArr = formData.skills.split(',').map(e => e.trim())

    const postData : PostData = { 
      name: formData.name,
      email: formData.email,
      description: formData.description,
      skills: tagArr,
    }
    try {
      const response: AxiosResponse<string> = await axiosBase.post<string>(`/project`, postData);
      const id = response.data;
    } catch (err) {
      // setSubmitError("An error occurred")
      console.log(postData)
      console.log("hello error")
    }
  }
  return (
  <div>
    <h2> Profile Settings </h2>
    <hr/>
    <div className="form-container">

      <label className='formLabel'>Name:</label>
      <input className='text-input' type="text"  name="name" value={formData.name} onChange={handleChange}/>
        {/* <div className='errorMessage'>
            {formFieldError.nameErr}
        </div> */}
      <label className='formLabel'>Email:</label>
      <input className='text-input' type="text"  name="email" value={formData.email} onChange={handleChange}/>

      <label className='formLabel'>About me:</label>
      <textarea className='text-textarea' name="description" value={formData.description} onChange={handleChange}/>

      <label className='formLabel'>Skills:</label>
      <input className='text-input' type="text"  name="skills" value={formData.skills} onChange={handleChange}/>
    </div>

    <div className='buttonContainer'>
      <button type="submit" onClick={onSubmit}>Save</button>
    </div>
  </div>
  )
}

export default PrivateProfile