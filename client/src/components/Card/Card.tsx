import React from 'react';

interface CardProps {
    title: string;
    body: string;
    link: string;
    linkDescription: string;
    svgTag?: JSX.Element;
}

const Card: React.FC<CardProps> = ({title, body, link, linkDescription, svgTag}) => {
    return (
        <div
            className="min-h-full flex flex-col w-1/4 mx-2 p-6 bg-white border border-gray-200 rounded-lg shadow dark:bg-white dark:border-gray-200">
            {svgTag}
            <a href={link}>
                <h5 className="mb-2 text-2xl font-semibold tracking-tight text-black dark:text-black">
                    {title}
                </h5>
            </a>
            <p className="mb-3 font-normal text-gray-900 dark:text-gray-900">
                {body}
            </p>
            <a href={link} className="mt-auto inline-flex justify-end items-center text-blue-600 hover:underline">
                {linkDescription}
                <svg className="w-5 h-5 ml-2" fill="currentColor" viewBox="0 0 20 20"
                     xmlns="http://www.w3.org/2000/svg">
                    <path
                        d="M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z"></path>
                    <path
                        d="M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z"></path>
                </svg>
            </a>
        </div>
    )
}

export default Card;