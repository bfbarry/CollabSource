import React from 'react';
import './NavBar.css';
import { useContext } from 'react';
import SignUpButton from './Components/SignUpButton';
import NavBarButton from './Components/NavBarButton';
import LogInButton from './Components/LogInButton';
import { AuthContext } from '../../context/AuthContext';
import SignOutButton from './Components/SignOutButton';

const NavBar: React.FC = () => {

    const {loggedIn} = useContext(AuthContext)

    return(
        <div id="nav-bar">
            <div id="left-nav-items">
                <a id="title" href="/">CollabSource</a>
                <NavBarButton text="Explore" pathToPage="/explore"/>
                <NavBarButton text="About" pathToPage=""/>
                <NavBarButton text="Other" pathToPage=""/>
            </div>
            {!loggedIn ?
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