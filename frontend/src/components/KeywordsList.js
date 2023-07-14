import axios from 'axios';
import React from 'react';
import { useState, useEffect } from 'react';
import KeywordElement from './KeywordElement';
import InputField from './InputField';
import { Store } from 'react-notifications-component';

function KeywordsList({ apiUrl }) {
  const [keywords, setKeywords] = useState([]);
  useEffect(() => {
    fetchKeywords();
  }, []);
  const fetchKeywords = async () => {
    try {
      const response = await axios.get(`${apiUrl}v1/keyword/get`);
      setKeywords(response.data);
      console.log(response.data.length);
    } catch (error) {
      console.error('Ошибка при получении списка групп:', error);
    }
  };

  const deleteKeywordLocally = (elem) => {
    setKeywords(keywords.filter((keyword) => keyword != elem));
  };

  const addKeyword = async (value) => {
    {
      try {
        const response = await axios.post(`${apiUrl}v1/keyword/add`, {
          keyword: value,
        });
        fetchKeywords();
        Store.addNotification({
          title: 'Успешно!',
          message: `Ключевое слово ${value} добавлено`,
          type: 'success',
          insert: 'top',
          container: 'top-right',
          animationIn: ['animate__animated', 'animate__fadeIn'],
          animationOut: ['animate__animated', 'animate__fadeOut'],
          dismiss: {
            duration: 3000,
            onScreen: true,
          },
        });
      } catch (error) {
        Store.addNotification({
          title: 'Ошибка!',
          message: `Ошибка при ключевого слова ${value}`,
          type: 'danger',
          insert: 'top',
          container: 'top-right',
          animationIn: ['animate__animated', 'animate__fadeIn'],
          animationOut: ['animate__animated', 'animate__fadeOut'],
          dismiss: {
            duration: 3000,
            onScreen: true,
          },
        });
        console.log(error);
      }
    }
  };

  return (
    <div className="card">
      <div className="card-body">
        <h4 className="card-title">Ключевые слова </h4>
        <h6 className="card-subtitle">всего {keywords ? keywords.length : 0}</h6>
        <InputField placeholder={'любое слово или фраза'} onClick={addKeyword} />
        <ul className="keywords">
          {keywords && keywords.length > 0 ? (
            keywords.map((keyword) => (
              <KeywordElement
                apiUrl={apiUrl}
                key={keyword}
                keyword={keyword}
                onDeleteLocally={deleteKeywordLocally}
              />
            ))
          ) : (
            <p>Нет ключевых слов</p>
          )}
        </ul>
      </div>
    </div>
  );
}
export default KeywordsList;
