import React from 'react';
import NavBarHeader from "../components/NavBar/NavBar";
import Hero from "../components/Hero/Hero";
import Benefits from "../components/Benefits/Benefits";
import Footer from "../components/Footer/Footer";

function AboutPage() {
  return (
    <div className="aboutpage-container min-h-screen flex flex-col">
      <NavBarHeader/>
      <div>
        Lorem Ipsum
      </div>
      <div className={"mt-auto"}>
        <Footer/>
      </div>
    </div>
  )
}

export default AboutPage;