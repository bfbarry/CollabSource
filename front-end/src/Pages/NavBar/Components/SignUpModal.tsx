import axios, { AxiosError } from 'axios';
import React, { ChangeEvent, Dispatch, SetStateAction, useState } from 'react';
import axiosBase from '../../../config/axiosConfig'
import './SignUpButtonAndModal.css'


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

    const modalStyle: React.CSSProperties = {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '30px',
        width: '400px',
        height: '350px',
        maxWidth: '80%',
        maxHeight: '80%',
        overflow: 'auto',
        display: 'block'
      };

      const modalOverlayStyle: React.CSSProperties = {
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100%',
        height: '100%',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'
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
    <div style={modalOverlayStyle}>
       <div style={modalStyle}>
       <button id="sign-up-modal-close-button" onClick={() => setShowSignUp(false)}>X</button>
        <div id="sign-up-modal-container">
            <p>Sign Up for CollabSource</p>
            <div id="sign-up-modal-input-container">
                <input style={getInputStyle(missingFields.name.color)} type="text" name="name" placeholder="name" value={formData.name} onChange={handleChange}/>
                <div className='SignUpModalError'>{missingFields.name.message}</div>
                <input style={getInputStyle(missingFields.email.color)} type="text"  name="email" placeholder="email" value={formData.email} onChange={handleChange}/>
                <div className='SignUpModalError'>{missingFields.email.message}</div>
                <input style={getInputStyle(missingFields.description.color)} type="text" name="description" placeholder="description (2-3 sentences)" value={formData.description} onChange={handleChange}/>
                <div className='SignUpModalError'>{missingFields.description.message}</div>
                <input style={getInputStyle(missingFields.password.color)} type="text" name="password" placeholder="password" value={formData.password} onChange={handleChange}/>
                <div className='SignUpModalError'>{missingFields.password.message}</div>
                <button type="submit" onClick={register}>Sign Up</button>
                <div className='SignUpModalError'>{signUpError}</div>
            </div>
        </div>

       </div>
    </div>
    );
}

export default SignUpModal;