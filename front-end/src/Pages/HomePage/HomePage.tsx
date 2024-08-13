import React, { useEffect, useState } from 'react';
import './HomePage.css';
import ProjectTile from './Components/ProjectTiles/ProjectTile';
import CreateProjectTile from './Components/ProjectTiles/CreateProjectTile';
import ScrollingTextBanner from '../ScrollingTextBanner/ScrollingTextBanner';
import SearchSection from './Components/SearchSection';
import axiosBase from '../../config/axiosConfig'
import { ProjectWId } from '../../types/project';


const HomePage: React.FC = () => {
    
    const [projects, setProjects] = useState<ProjectWId[]>([]);

    useEffect(() => {
        axiosBase.get('/projects?page=1&size=3')
        .then(response => {
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
                    <ProjectTile 
                    key={value._id}
                    _id={value._id}
                    name={value.name} 
                    description={value.description} 
                    category={value.category} 
                    tags={value.tags} 
                    seeking={value.seeking}
                    />
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