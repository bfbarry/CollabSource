import React from 'react';
import axios from 'axios';
import './NavBar.css';

interface Props {
    text: string;
    pathToPage: string;
}

const NavBarButton: React.FC<Props> = ({text, pathToPage}) => { 

    const login = async () => {
        try {
        const response = await axios.post(`http://localhost:8000/auth/login`, { email: "test1", password : "1232" });
        console.log(response.data.token)
        } catch (error){
            console.log(error)
        }
    } 

    return(
        <button onClick={login}>{text}</button>
    );
}

export default NavBarButton;