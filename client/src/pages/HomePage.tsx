import React from 'react';
import NavBarHeader from "../components/NavBar/NavBar";
import Hero from "../components/Hero/Hero";

function HomePage() {
  return (
    <div className="homepage-container">
      <NavBarHeader/>
      <Hero/>
    </div>
  );
};

export default HomePage;
