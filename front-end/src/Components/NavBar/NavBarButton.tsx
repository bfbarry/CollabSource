import React from 'react';
import './NavBar.css';

interface Props {
    text: string;
    pathToPage: string;
}

const NavBarButton: React.FC<Props> = ({text, pathToPage}) => { 
    return(
        <a href="">{text}</a>
    );
}

export default NavBarButton;