import React from 'react';
import Card from "../Card/Card";

const Benefits = () => {
    return (
        <section className="my-6 flex flex-col w-full justify-items-center">
            <div className="font-bold text-2xl text-center">Here's how decider can help</div>
            <div className="mt-4 flex w-3/4 self-center">
                <Card title={"Create polls"}
                      body={"Come to a consensus on decisions quickly. Post something and have others vote on the options."}
                      link={"www.google.com"}
                      linkDescription="View our demo"
                      svgTag={
                          <svg className="w-10 h-10 mb-2 text-primary dark:text-primary" aria-hidden="true" fill="none"
                               stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24"
                               xmlns="http://www.w3.org/2000/svg">
                              <path
                                  d="M16.5 6v.75m0 3v.75m0 3v.75m0 3V18m-9-5.25h5.25M7.5 15h3M3.375 5.25c-.621
                                  0-1.125.504-1.125 1.125v3.026a2.999 2.999 0 010 5.198v3.026c0 .621.504 1.125
                                  1.125 1.125h17.25c.621 0 1.125-.504 1.125-1.125v-3.026a2.999 2.999 0
                                  010-5.198V6.375c0-.621-.504-1.125-1.125-1.125H3.375z"
                                  stroke-linecap="round"
                                  stroke-linejoin="round">
                              </path>
                          </svg>
                      }
                />

                <Card title={"Share suggestions"}
                      body={"Post suggestions and ideas to keep track and prioritise improvements."}
                      link={"www.google.com"}
                      linkDescription="View our demo"
                      svgTag={
                          <svg className="w-10 h-10 mb-2 text-primary dark:text-primary" aria-hidden="true" fill="none"
                               stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24"
                               xmlns="http://www.w3.org/2000/svg">
                              <path
                                  d="M21.75 9v.906a2.25 2.25 0 01-1.183 1.981l-6.478 3.488M2.25 9v.906a2.25 2.25
                                  0 001.183 1.981l6.478 3.488m8.839 2.51l-4.66-2.51m0 0l-1.023-.55a2.25 2.25 0
                                  00-2.134 0l-1.022.55m0 0l-4.661 2.51m16.5 1.615a2.25 2.25 0 01-2.25 2.25h-15a2.25
                                  2.25 0 01-2.25-2.25V8.844a2.25 2.25 0 011.183-1.98l7.5-4.04a2.25 2.25 0 012.134
                                  0l7.5 4.04a2.25 2.25 0 011.183 1.98V19.5z"
                                  stroke-linecap="round"
                                  stroke-linejoin="round">
                              </path>
                          </svg>
                      }
                />

                <Card title={"Random selection"}
                      body={"Can’t decide? Make use of random selection via a “Spin the Wheel” or other randomisers to help make decisions for you."}
                      link={"www.google.com"}
                      linkDescription="View our demo"
                      svgTag={
                          <svg className="w-10 h-10 mb-2 text-primary dark:text-primary" aria-hidden="true" fill="none"
                               stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24"
                               xmlns="http://www.w3.org/2000/svg">
                              <path
                                  d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0
                                  3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 11-18 0 9
                                  9 0 0118 0zm-9 5.25h.008v.008H12v-.008z"
                                  stroke-linecap="round"
                                  stroke-linejoin="round">
                              </path>
                          </svg>
                      }
                />

                <Card title={"Organise events"}
                      body={"Don’t let events stay online only! Create and share calendar invites to others."}
                      link={"www.google.com"}
                      linkDescription="View our demo"
                      svgTag={
                          <svg className="w-10 h-10 mb-2 text-primary dark:text-primary" aria-hidden="true" fill="none"
                               stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24"
                               xmlns="http://www.w3.org/2000/svg">
                              <path
                                  d="M6.75 3v2.25M17.25 3v2.25M3 18.75V7.5a2.25 2.25 0 012.25-2.25h13.5A2.25 2.25 0 0121
                                   7.5v11.25m-18 0A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75m-18 0v-7.5A2.25
                                   2.25 0 015.25 9h13.5A2.25 2.25 0 0121 11.25v7.5"
                                  stroke-linecap="round"
                                  stroke-linejoin="round">
                              </path>
                          </svg>
                      }
                />
            </div>
        </section>
    )
}

export default Benefits;

