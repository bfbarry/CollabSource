import axios, { AxiosError } from 'axios';
import React, { ChangeEvent, Dispatch, SetStateAction, useState } from 'react';
import axiosBase from '../../config/axiosConfig'
import './Modal.css'


interface Props {
    setShowSignUp: Dispatch<SetStateAction<Boolean>>
}

interface UserRequestBody {
    name: string;
    password: string;
    email: string;
    description: string;
}

interface MissingFieldType{
    name: {
        message: string,
        color: string
    };
    password: {
        message: string,
        color: string
    };
    email: {
        message: string,
        color: string
    };
    description: {
        message: string,
        color: string
    };
}

const SignUpModal: React.FC<Props> = ({setShowSignUp}) => { 
    const [ signUpError, setSignUpError ] = useState<String>("");
    const [ formData, setFormData ] = useState<UserRequestBody>({
        name: '',
        password: '',
        email: '',
        description: '',
    })

    const [ missingFields, setMissingFields ] = useState<MissingFieldType>({
        name: { message:'', color:''},
        password: { message:'', color:''},
        email: { message:'', color:''},
        description: { message:'', color:''}
    })

    const register = async () => {
        if (formData.email === "" || formData.name === "" || formData.password === "" || formData.description === "") { 
            setMissingFields({
                name: !formData.name ? { message:'Missing name', color:'red'} : { message:'', color:''},
                email: !formData.email ? { message:'Missing email', color:'red'} : { message:'', color:''},
                description: !formData.description ? { message:'Missing description', color:'red'} : { message:'', color:''},
                password: !formData.password ? { message:'Missing password', color:'red'} : { message:'', color:''}
            })
            return;
        }

        try {
        const response = await axiosBase.post(`/auth/register`, 
        { email: formData.email, password : formData.password, description: formData.description, name: formData.name });
        setShowSignUp(false);
        } catch (error){
            if (axios.isAxiosError(error)) {
                const axiosError = error as AxiosError;
                const statusCode = axiosError.response?.status;
                if (statusCode == 422) {
                    setSignUpError("User already exists");
                } else if (statusCode == 500) {
                    setSignUpError("Unable to register user");
                } else {
                    setSignUpError("Something went wrong");
                }
            }   
        }
    } 

    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({           
            ...formData,
            [name]: value,
        });
    };

      const getInputStyle = (color:string): React.CSSProperties => {
        return {
          borderColor: color,
          borderRadius: '10px',
          border: 'solid 1px',
          padding: '5px'
        };
      };

    return(
    <div className="modalBackdrop">
       <div className="modalStyleParent modalStyleSignup">
       <button className="closeButton" onClick={() => setShowSignUp(false)}>X</button>
        <div id="sign-up-modal-container">
            <p>Sign Up for CollabSource</p>
            <div id="sign-up-modal-input-container">
                <input style={getInputStyle(missingFields.name.color)} type="text" name="name" placeholder="name" value={formData.name} onChange={handleChange}/>
                <div className='errorMessage'>{missingFields.name.message}</div>
                <input style={getInputStyle(missingFields.email.color)} type="text"  name="email" placeholder="email" value={formData.email} onChange={handleChange}/>
                <div className='errorMessage'>{missingFields.email.message}</div>
                <input style={getInputStyle(missingFields.description.color)} type="text" name="description" placeholder="description (2-3 sentences)" value={formData.description} onChange={handleChange}/>
                <div className='errorMessage'>{missingFields.description.message}</div>
                <input style={getInputStyle(missingFields.password.color)} type="text" name="password" placeholder="password" value={formData.password} onChange={handleChange}/>
                <div className='errorMessage'>{missingFields.password.message}</div>
                <button type="submit" onClick={register}>Sign Up</button>
                <div className='errorMessage'>{signUpError}</div>
            </div>
        </div>

       </div>
    </div>
    );
}

export default SignUpModal;