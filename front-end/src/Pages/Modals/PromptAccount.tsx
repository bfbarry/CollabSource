import React, { ChangeEvent, Dispatch, SetStateAction, useContext, useState } from 'react';
import axiosBase from '../../config/axiosConfig'
import { AuthContext } from '../../context/AuthContext';
import LogInButton from '../NavBar/Components/LogInButton';
import SignUpButton from '../NavBar/Components/SignUpButton';
import './Modal.css'


interface Props {
  setShowPromptAccount: Dispatch<SetStateAction<Boolean>>
  setShowLogIn: Dispatch<SetStateAction<Boolean>>
  setShowSignUp: Dispatch<SetStateAction<Boolean>>
}


const PromptAccount: React.FC<Props> = ({setShowPromptAccount, setShowLogIn, setShowSignUp} )=> {
  const onLoginClick = () => {
    setShowLogIn(true)
    setShowPromptAccount(false)
  }

  const onSignupClick = () => {
    setShowSignUp(true)
    setShowPromptAccount(false)
  }
  return(
  <div className="modalBackdrop">
    <div className="modalStyleParent modalStyleLogin">
      <button className="closeButton" onClick={() => setShowPromptAccount(false)}>X</button>
      <div className='buttonParent'>
        
        <p>To create a project, please log in: </p>
        <button className="promptButton login" onClick={onLoginClick}>Log In</button>

        <p>Don't have an account? Sign up here:</p>
        <button className="promptButton signUp" onClick={onSignupClick}>Sign Up</button>
      </div>

    </div>
  </div>
  );
}

export default PromptAccount;

