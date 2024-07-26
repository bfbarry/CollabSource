import React, { useState } from 'react';
import SignUpModal from './SignUpModal';

const SignUpButton: React.FC = () => { 

    const [showSignUp, setShowSignUp ] = useState<Boolean>(false);

    return(
        <div>
            {
                showSignUp && <SignUpModal setShowSignUp={setShowSignUp}/>
            }
            <button onClick={() => setShowSignUp(true)}>Sign Up</button>
        </div>
    );
}

export default SignUpButton;