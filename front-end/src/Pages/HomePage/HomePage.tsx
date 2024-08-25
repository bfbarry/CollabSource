import React, { useEffect, useState } from 'react';
import './HomePage.css';
import ScrollingTextBanner from './Components/ScrollingTextBanner';
import axiosBase from '../../config/axiosConfig'
import { ProjectWId } from '../../types/project';
import CreateProjectTile from '../Common/ProjectTiles/CreateProjectTile';
import ProjectTile from '../Common/ProjectTiles/ProjectTile';
import { Filters } from "../../types/project";
import SearchBar from '../Common/SearchBar';

const HomePage: React.FC = () => {
  
  const [projects, setProjects] = useState<ProjectWId[]>([]);
  const [searchQuery, setSearchQuery] = useState<string>('')

  const moreText = 'Explore more projects >'
  useEffect(() => {
    const filters: Filters = {categories:[], searchQuery:''}
    axiosBase.post('/projects?page=1&size=3', filters)
    .then(response => {
      setProjects(response.data.items)
    })
    .catch(error => {
      console.log(error)
    })
    return 
    }, []);

  return(
    <div>  
      <div className="search-section">
        <h2>Find collaborators for your passion projects.</h2>
        <SearchBar 
          searchParams={new URLSearchParams()} 
          setSearchParams={()=>{}}
          searchQuery={searchQuery}
          setSearchQuery={setSearchQuery}
          redirect={true}/>
      </div>
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
        <div>
          <a className='more-projects-link' href='/explore'>{moreText}</a>
        </div>
      </div>
      <div id="scroll-container">
        <ScrollingTextBanner/>
      </div>
    </div>
  );
}

export default HomePage;