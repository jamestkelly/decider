import React from 'react';
import './BoardButton.css';

interface BoardButtonProps {
  icon?: React.ReactNode;
  imgSrc?: string;
  text: string;
  onClick: () => void;
}

const BoardButton: React.FC<BoardButtonProps> = ({ icon, imgSrc, text, onClick }) => {
  return (
    <div className="board-button group" onClick={onClick}>
      { icon ? icon : imgSrc ? <img src={imgSrc} alt={text} /> : null }
      <span className="board-button-tooltip group-hover:scale-100">{ text }</span>
    </div>
  )
}

export default BoardButton;