import React, { ChangeEvent, Dispatch, SetStateAction, useContext, useState } from 'react';
import axiosBase from '../../../config/axiosConfig'
import { AuthContext } from '../../../context/AuthContext';
import './LogInButtonAndModal.css'
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
        const response = await axiosBase.post(`/auth/login`, { email: formData.email, password : formData.password });
        // console.log(response.data.token)
        const user = { userID: formData.email, token: response.data.token, loggedIn: true}
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

    const modalStyle: React.CSSProperties = {
        position: 'absolute',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: '#fff',
        padding: '20px',
        borderRadius: '30px',
        width: '400px',
        height: '250px',
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

    return(
    <div style={modalOverlayStyle}>
        <div style={modalStyle}>
            <div id = "LogInModal-container">
                <div>
                    <button id="LogInModalCloseButton" onClick={() => SetShowLogIn(false)}>X</button>
                    <p>Sign In to CollabSource</p>
                </div>
                <div id='LogInModal-input-container'>
                    <input type="text"  name="email" placeholder="email" value={formData.email} onChange={handleChange}/>
                    <div className='LogInModalMissingFieldMessage'>
                        {formFieldError.emailError}
                    </div>
                
                    <input type="text" name="password" placeholder="password" value={formData.password} onChange={handleChange}/>
                    <div className='LogInModalMissingFieldMessage'>
                        {formFieldError.passwordError}
                    </div>
                    <button type="submit" onClick={login}>Sign In</button>
                    <div className='LogInModalMissingFieldMessage'>
                        {logInError}
                    </div>
                </div>
            </div>
        </div>
    </div>
    );
}

export default LogInModal;

