import React, { useState } from 'react';
import SignUpModal from '../../Modals/SignUpModal';
import './buttons.css'; 
import { createPortal } from 'react-dom';

const SignUpButton: React.FC = () => { 

  const [showSignUp, setShowSignUp ] = useState<Boolean>(false);

  return(
    <div>
      {
        showSignUp && 
        <>
        {createPortal(
        <SignUpModal setShowSignUp={setShowSignUp}/>        ,
        document.body
        )
        }
        </>
        
      }
      <button id="sign-up-button" onClick={() => setShowSignUp(true)}>Sign Up</button>
    </div>
  );
}

export default SignUpButton;