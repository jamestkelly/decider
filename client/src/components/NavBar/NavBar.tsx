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
          <img src={logo} alt="Decider" style={{ maxHeight: '50px', maxWidth: '50px' }}/>
          <div className="text-lg font-bold m-2.5">Decider</div>
        </div>
        <div className="hidden md:flex space-x-12 items-center">
          <Button text={"About"} className="hover:text-selected-text transition-all duration-150 ease-linear" onClick={scrollTop}/>
          <Button text={"Support"} className="hover:text-selected-text transition-all duration-150 ease-linear" onClick={scrollTop}/>
          <Button text={"Documentation"} className="hover:text-selected-text transition-all duration-150 ease-linear" onClick={scrollTop}/>
          <button
            className="relative inline-block text-lg group"
            onClick={handleOpenModal}
          >
            <span className="relative z-10 block px-5 py-3 overflow-hidden font-medium leading-tight text-white transition-colors duration-300 ease-out bg-theme rounded group-hover:text-black">
              <span className="absolute inset-0 w-full h-full px-5 py-3 rounded-lg bg-theme"></span>
              <span className="absolute left-0 w-48 h-48 -ml-2 transition-all duration-300 origin-top-right -rotate-90 -translate-x-full translate-y-12 bg-white group-hover:-rotate-180 ease"></span>
              <span className="relative">Login</span>
            </span>
            <span
              className="absolute bottom-0 right-0 w-full h-12 -mb-1 -mr-1 transition-all duration-200 ease-linear bg-gray-900 rounded-lg group-hover:mb-0 group-hover:mr-0"
              data-rounded="rounded-lg"
            ></span>
          </button>
        </div>
      </div>
    </header>
  );
};

export default NavBarHeader;