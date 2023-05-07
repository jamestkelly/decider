import React from 'react';
import Button from "../Buttons/Button/Button";

const Hero = () => {
  return (
    <section className="bg-white mt-2">
      <div
        className="bg-primary rounded-2xl grid max-w-screen-xl w-5/6 px-8 py-8 mx-auto lg:gap-8 xl:gap-0 lg:py-16 lg:grid-cols-12">
        <div className="mr-auto place-self-center lg:col-span-7">
          <h1
            className="max-w-2xl mb-4 text-4xl font-extrabold tracking-tight leading-none md:text-5xl xl:text-6xl dark:text-white">
            Make decisions easier
          </h1>
          <p className="max-w-2xl mb-6 font-light text-white lg:mb-8 md:text-lg lg:text-xl dark:text-white">
            Come to a consensus faster, share suggestions with ease, and use the power of random selection
            to make decisions, with a powerful and easy-to-use tool.
          </p>
          <div className="flex flex-1 justify-center">
            <Button text={"Get started"}
                    className={"bg-white text-primary font-semibold rounded px-10 py-2.5 hover:bg-tertiary transition-all duration-150 ease-linear hover:text-white"}/>
          </div>
        </div>
        <div className="hidden lg:mt-0 lg:col-span-5 lg:flex">
          <img src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/hero/phone-mockup.png"
               alt="mockup"/>
        </div>
      </div>
    </section>
  )
}

export default Hero;