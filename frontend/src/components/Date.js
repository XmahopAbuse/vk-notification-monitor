import React from 'react';

function DateComponent({ date }) {
  const formattedDate = new Date(date).toLocaleString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric',
    hour: 'numeric',
    minute: 'numeric',
    second: 'numeric',
  });

  return (
    <div>
      <p>{formattedDate}</p>
    </div>
  );
}

export default DateComponent;
