import React, { useState } from 'react';
import LogInModal from '../../Modals/LogInModal';
import "./buttons.css"
import { createPortal } from 'react-dom';


const LogInButton: React.FC = () => { 

  const [showLogIn, setShowLogIn] = useState<Boolean>(false);

  return(
    <div>
      {
        showLogIn && 
        <>
          {createPortal(
          <LogInModal SetShowLogIn={setShowLogIn}/>,
          document.body
          )
          }
        </>
          
      }
      <button id="logInButton" onClick={() => setShowLogIn(true)}>Sign In</button>
    </div>
  );
}

export default LogInButton;

