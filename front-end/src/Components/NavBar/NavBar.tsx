import React, { useState } from 'react';
import './NavBar.css';
import NavBarButton from './NavBarButton';

const NavBar: React.FC = () => {
    return(
        <div id="nav-bar">
            <div id="left-nav-items">
                <a id="title" href="">CollabSource</a>
                <NavBarButton text="Explore" pathToPage=""/>
                <NavBarButton text="About" pathToPage=""/>
                <NavBarButton text="Other" pathToPage=""/>
            </div>
            <div id="right-nav-items">
                <NavBarButton text="Log In" pathToPage=""/>
                <NavBarButton text="Sign Up" pathToPage=""/>
            </div>
        </div>
    );
}

export default NavBar;