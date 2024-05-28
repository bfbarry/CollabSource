import React from 'react';
import './HomePage.css';
import ProjectTile from './Components/ProjectTile/ProjectTile';
import ScrollingTextBanner from '../Components/ScrollingTextBanner/ScrollingTextBanner';
import SearchSection from './Components/SearchSection';

const HomePage: React.FC = () => {
    return(
    <>    
    {/* K TODO make component */}
        <SearchSection/>
        <div id="explore-projects-section">
            <h2>Explore Open Projects</h2>
            <div id="project-tiles-section">
                <ProjectTile/>
                <ProjectTile/>
                <ProjectTile/>
            </div>
        </div>
        <div id="scroll-container">
            <ScrollingTextBanner/>
        </div>
    </>
    );
}

export default HomePage;