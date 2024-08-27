import React from 'react';
import './ScrollingTextBanner.css';

const ScrollingTextBanner: React.FC = () => {
  const bannerItems = ["Graphic Design", 'Digital Marketing', 'Software Engineering', '3D Printing', 'UX Design', 'Art',
     'Environmental Initiatives', 'Computer Science', 'Education', 'Outdoor Activities', 'Music', 'Dance'
  ]
  return(
    <div className="wrapper">
      <div className="marquee">
        <p>
        {bannerItems.map((val) => (
          <span key={val}>{val + '   -   '} </span>
        ))}
        </p>
        <p>
        {bannerItems.map((val) => (
          <span key={val}>{val + '   -   '} </span>
        ))}
        </p>
      </div>
    </div>
  )
}

export default ScrollingTextBanner;