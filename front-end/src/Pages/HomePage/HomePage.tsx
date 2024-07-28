import React, { useEffect, useState } from 'react';
import './HomePage.css';
import ProjectTile from './Components/ProjectTiles/ProjectTile';
import CreateProjectTile from './Components/ProjectTiles/CreateProjectTile';
import ScrollingTextBanner from '../ScrollingTextBanner/ScrollingTextBanner';
import SearchSection from './Components/SearchSection';
import axiosBase from '../../config/axiosConfig'

interface Project {
    id: string;
    name       : string;
    description: string;
    category   : string;
    tags       : string[];
}

const HomePage: React.FC = () => {
    
    const [projects, setProjects] = useState<Project[]>([]);

    useEffect(() => {
        axiosBase.get('/projects?page=1&size=3')
        .then(response => {
            console.log(response)
            setProjects(response.data.data)
        })
        .catch(error => {
            console.log(error)
        })
        return 
      }, []);

    return(
        <div>    
            <SearchSection/>
            <div id="explore-projects-section">
                <div id="project-tiles-section">
                <CreateProjectTile/>
                {projects.map((value) => (
                    <ProjectTile name={value.name} 
                    description={value.description} 
                    category={value.category} 
                    tags={value.tags} 
                    id={value.id}/>
                ))}
                </div>
            </div>
            <div id="scroll-container">
                <ScrollingTextBanner/>
            </div>
        </div>
    );
}

export default HomePage;