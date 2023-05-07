import React from 'react';
import NavBarHeader from "../components/NavBar/NavBar";
import Hero from "../components/Hero/Hero";
import Footer from "../components/Footer/Footer";
import Benefits from "../components/Benefits/Benefits";
import Carousel from "../components/Carousel/Carousel";

function HomePage() {
  return (
    <div className="homepage-container min-h-screen flex flex-col">
      <NavBarHeader/>
      <div>
        <Hero/>
        <Benefits/>
        <Carousel/>
      </div>
      <div className={"mt-auto"}>
        <Footer/>
      </div>
    </div>
  );
}

export default HomePage;
