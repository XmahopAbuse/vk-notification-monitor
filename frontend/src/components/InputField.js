import React from 'react';
import { useState } from 'react';
import { Store } from 'react-notifications-component';

function InputField({ placeholder, onClick }) {
  const [inputValue, setInputValue] = useState('');
  const onButtonClick = (e, value) => {
    e.preventDefault();
    setInputValue('');
    if (inputValue == '') {
      Store.addNotification({
        title: 'Ошибка',
        message: `Невозможно добавить пустую строку`,
        type: 'danger',
        insert: 'top',
        container: 'top-right',
        animationIn: ['animate__animated', 'animate__fadeIn'],
        animationOut: ['animate__animated', 'animate__fadeOut'],
        dismiss: {
          duration: 1000,
          onScreen: true,
        },
      });
    } else {
      onClick(inputValue);
    }
  };

  const onInputChange = (event) => {
    setInputValue(event.target.value);
  };
  return (
    <div className="input" style={{ marginBottom: '10px' }}>
      <form>
        <input
          value={inputValue}
          onChange={onInputChange}
          placeholder={placeholder}
          style={{ width: '94%' }}
          type="text"></input>
        <button onClick={onButtonClick} style={{ width: '5%', marginLeft: '1%' }}>
          +
        </button>
      </form>
    </div>
  );
}

export default InputField;
