import React, { useState } from 'react';
import SignUpModal from '../../Modals/SignUpModal';
import './buttons.css'; 

const SignUpButton: React.FC = () => { 

    const [showSignUp, setShowSignUp ] = useState<Boolean>(false);

    return(
        <div>
            {
                showSignUp && <SignUpModal setShowSignUp={setShowSignUp}/>
            }
            <button id="sign-up-button" onClick={() => setShowSignUp(true)}>Sign Up</button>
        </div>
    );
}

export default SignUpButton;