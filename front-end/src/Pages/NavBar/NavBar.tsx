import React from 'react';
import './NavBar.css';
import { useContext } from 'react';
import SignUpButton from './Components/SignUpButton';
import NavBarButton from './Components/NavBarButton';
import LogInButton from './Components/LogInButton';
import { SignedInContext } from '../../context/SignedInContext';
import SignOutButton from './Components/SignOutButton';

const NavBar: React.FC = (setLogedInUser) => {

    const signedIn = useContext(SignedInContext)

    return(
        <div id="nav-bar">
            <div id="left-nav-items">
                <a id="title" href="/">CollabSource</a>
                <NavBarButton text="Explore" pathToPage="/explore"/>
                <NavBarButton text="About" pathToPage=""/>
                <NavBarButton text="Other" pathToPage=""/>
            </div>
            {!signedIn.signedInUser ?
                <div id="right-nav-items">
                    <LogInButton/>
                    <SignUpButton/>
                </div>
                :
                <div id="right-nav-items">
                    <SignOutButton/>
                </div>
            }
        </div>
    );
}

export default NavBar;