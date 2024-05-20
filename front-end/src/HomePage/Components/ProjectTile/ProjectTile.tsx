import React from 'react';
import './ProjectTile.css';


const ProjectTile: React.FC = () => {
    return(
        <div id="project-tile">
            <h3>3D Printing</h3>
            <p id="project-tile-description">Description Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>                <p id="project-tile-description"><b>Seeking:</b> Position1, Position2</p>
         </div>
    );
}

export default ProjectTile;