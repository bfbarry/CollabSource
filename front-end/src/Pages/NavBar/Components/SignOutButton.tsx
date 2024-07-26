import React, { useContext } from 'react';
import { SignedInContext } from '../../../context/SignedInContext';

const SignOutButton: React.FC = () => { 

    const signedIn = useContext(SignedInContext)

    const signOut = () => {
        localStorage.removeItem("access_token");
        signedIn.setSignedInUser(false);
    }

    return(
        <div>
            <button onClick={() => signOut()}>Sign Out</button>
        </div>
    );
}

export default SignOutButton;