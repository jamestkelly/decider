import React from 'react';
import NavBarHeader from "../components/NavBar/NavBar";
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