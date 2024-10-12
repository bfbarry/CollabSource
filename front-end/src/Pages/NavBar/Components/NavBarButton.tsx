import React from 'react';
import '../NavBar.css';
import { NavLink } from 'react-router-dom';

interface Props {
    text: string;
    pathToPage: string;
}

const NavBarButton: React.FC<Props> = ({text, pathToPage}) => { 
    return(
        <NavLink to={pathToPage} className="NavLink">{text}</NavLink>
    );
}

export default NavBarButton;