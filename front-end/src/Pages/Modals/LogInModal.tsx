import React, { ChangeEvent, Dispatch, SetStateAction, useContext, useState } from 'react';
import axiosBase from '../../config/axiosConfig'
import { AuthContext } from '../../context/AuthContext';
import './Modal.css'
import useLogin from '../../hooks/useLogin';
interface Props {
    SetShowLogIn: Dispatch<SetStateAction<Boolean>>
}

interface UserLoginRequestBody {
    email: string,
    password: string
}

interface FormFieldsError {
    emailError: string,
    passwordError: string
}

interface AuthResponse {
    token: string,
    userId: string
}

const LogInModal: React.FC<Props> = ({SetShowLogIn} )=> {

    const { login, logInError} = useLogin()

    const [formData, setFormData ] = useState<UserLoginRequestBody>({
        password: '',
        email: ''
    })
    const [formFieldError, setFormFieldError] = useState<FormFieldsError>({emailError:"",passwordError:""});


    const onLoginClick = async () => {
        if (formData.email === "" || formData.password === "" ) {
            setFormFieldError({
                emailError: !formData.email ? "Missing email" : "" ,
                passwordError: !formData.password ? "Missing password" : "" 
            })
            return;
        } else {
            setFormFieldError({
                emailError: "" ,
                passwordError: "" 
            })
        }
        login(formData.email, formData.password)
        if (logInError === '') {
            SetShowLogIn(false)
        } 
    }

    const handleChange = (e: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setFormData({           
            ...formData,
            [name]: value,
        });
    };

    return(
    <div className="modalBackdrop">
        <div className="modalStyleParent modalStyleLogin">
            <div id = "LogInModal-container">
                <div>
                    <button className="closeButton" onClick={() => SetShowLogIn(false)}>X</button>
                    <p>Sign In to CollabSource</p>
                </div>
                <div id='LogInModal-input-container'>
                    <input type="text"  name="email" placeholder="email" value={formData.email} onChange={handleChange}/>
                    <div className='errorMessage'>
                        {formFieldError.emailError}
                    </div>
                
                    <input type="text" name="password" placeholder="password" value={formData.password} onChange={handleChange}/>
                    <div className='errorMessage'>
                        {formFieldError.passwordError}
                    </div>
                    <button type="submit" onClick={onLoginClick}>Sign In</button>
                    <div className='errorMessage'>
                        {logInError}
                    </div>
                </div>
            </div>
        </div>
    </div>
    );
}

export default LogInModal;

