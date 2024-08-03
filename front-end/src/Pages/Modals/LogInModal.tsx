import React, { ChangeEvent, Dispatch, SetStateAction, useContext, useState } from 'react';
import axiosBase from '../../config/axiosConfig'
import { AuthContext } from '../../context/AuthContext';
import './Modal.css'
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

    const { authDispatch } = useContext(AuthContext)

    const [formData, setFormData ] = useState<UserLoginRequestBody>({
        password: '',
        email: ''
    })
    const [formFieldError, setFormFieldError] = useState<FormFieldsError>({emailError:"",passwordError:""});

    const [logInError, setLogInError] = useState<String>("");

    const login = async () => {
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
        try {
        const response = await axiosBase.post<AuthResponse>(`/auth/login`, { email: formData.email, password : formData.password });
        // console.log(response.data.token)
        const user = { userID: response.data.userId, token: response.data.token, loggedIn: true}
        localStorage.setItem("auth_context_state", JSON.stringify(user));
        authDispatch({ type: 'LOG_IN', payload: user });
        // authContext.setEmail(formData.email);
        SetShowLogIn(false)
        } catch (error){
            setLogInError("Invalid username or password")
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
                    <button type="submit" onClick={login}>Sign In</button>
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

