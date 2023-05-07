import React from 'react';
import logo from '../../assets/images/decider_logo_raw.png';
import Button from "../Buttons/Button/Button";

const NavBarHeader = () => {
  const scrollTop = () => {
    // TODO: Add logic for scroll to top
  };

  const handleOpenModal = () => {
    // TODO: Add logic for opening modal
  };

  return (
    <header className="py-2 sticky top-0 z-50 bg-white">
      <div className="container flex justify-between items-center mx-auto px-8 md:px-14 lg:px-24 w-full">
        <div className="container flex align-middle">
          <img src={'https://firebasestorage.googleapis.com/v0/b/decider-c60e5.appspot.com/o/public-assets%2Fdecider_logo_raw.png?alt=media&token=dd948dc8-eaa2-4302-ba36-b0336ce77a8a'} alt="Decider" style={{maxHeight: '50px', maxWidth: '50px'}}/>
          <div className="text-lg font-bold m-2.5">Decider</div>
        </div>
        <div className="hidden md:flex space-x-12 items-center">
          <Button text={"About"}
                  className="font-semibold hover:text-selected-text transition-all duration-150 ease-linear"
                  onClick={scrollTop}/>
          <Button text={"Support"}
                  className="font-semibold hover:text-selected-text transition-all duration-150 ease-linear"
                  onClick={scrollTop}/>
          <Button text={"Documentation"}
                  className="font-semibold hover:text-selected-text transition-all duration-150 ease-linear"
                  onClick={scrollTop}/>
          <Button text={"Login"}
                  className="font-bold text-white rounded px-10 pt-2 pb-3 bg-primary hover:bg-tertiary transition-all duration-150 ease-linear"
                  onClick={scrollTop}/>
        </div>
      </div>
    </header>
  );
};

export default NavBarHeader;