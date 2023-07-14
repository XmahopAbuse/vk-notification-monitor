import React from 'react';
import axios from 'axios';
import { Store } from 'react-notifications-component';

function KeywordElement({ keyword, onDeleteLocally, apiUrl }) {
  const onClickKeyword = async () => {
    try {
      await axios.delete(`${apiUrl}v1/keyword/${keyword}/delete`);
      onDeleteLocally(keyword);
      Store.addNotification({
        title: 'Успешно!',
        message: `Ключевое слово "${keyword}" удалено`,
        type: 'success',
        insert: 'top',
        container: 'top-right',
        animationIn: ['animate__animated', 'animate__fadeIn'],
        animationOut: ['animate__animated', 'animate__fadeOut'],
        dismiss: {
          duration: 1000,
          onScreen: true,
        },
      });
    } catch {
      Store.addNotification({
        title: 'Ошибка!',
        message: `Ошибка при удалении ключевого слова`,
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
      console.log('err');
    }
  };

  return (
    <a onClick={onClickKeyword} href="#">
      <li>{keyword}</li>
      <span className="keywordClose">удалить</span>
    </a>
  );
}

export default KeywordElement;
