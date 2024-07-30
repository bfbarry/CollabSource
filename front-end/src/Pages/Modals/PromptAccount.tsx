import React, { ChangeEvent, Dispatch, SetStateAction, useContext, useState } from 'react';
import axiosBase from '../../config/axiosConfig'
import { AuthContext } from '../../context/AuthContext';
import LogInButton from '../NavBar/Components/LogInButton';
import SignUpButton from '../NavBar/Components/SignUpButton';
import './Modal.css'


interface Props {
  SetShowPromptAccount: Dispatch<SetStateAction<Boolean>>
  // SetShowLogIn: Dispatch<SetStateAction<Boolean>>
  // SetShowSignUp: Dispatch<SetStateAction<Boolean>>
}


const PromptAccount: React.FC<Props> = ({SetShowPromptAccount} )=> {
  return(
  <div className="modalBackdrop">
    <div className="modalStyleParent modalStyleLogin">
      <button className="closeButton" onClick={() => SetShowPromptAccount(false)}>X</button>
      <p>To create a project, please log in</p>
      <LogInButton />
      <p>Don't have an account? Sign up here:</p>
      <SignUpButton/>
    </div>
  </div>
  );
}

export default PromptAccount;

