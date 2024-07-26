import React, { useState } from 'react';
import LogInModal from './LogInModal';
import "./LogInButtonAndModal.css"

const LogInButton: React.FC = () => { 

    const [showLogIn, setShowLogIn] = useState<Boolean>(false);

    return(
        <div>
            {
                showLogIn && <LogInModal SetShowLogIn={setShowLogIn}/>
            }
            <button id="logInButton" onClick={() => setShowLogIn(true)}>Sign In</button>
        </div>
    );
}

export default LogInButton;

