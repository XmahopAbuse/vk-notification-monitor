import React, { useState, useEffect } from 'react';
import GroupElement from './GroupElement';
import axios from 'axios';
import InputField from './InputField';
import { Store } from 'react-notifications-component';

export default function GroupList({ apiUrl }) {
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    fetchGroups();
  }, []);

  const fetchGroups = async () => {
    try {
      const response = await axios.get(`${apiUrl}v1/group/get`);
      setGroups(response.data);
      console.log(response);
    } catch (error) {
      console.error('Ошибка при получении списка групп:', error);
    }
  };

  const deleteGroupLocally = (e) => {
    setGroups(groups.filter((group) => group.name != e.name));
  };

  const addGroup = async (value) => {
    {
      try {
        const response = await axios.post(`${apiUrl}v1/group/add`, {
          address: value,
        });
        Store.addNotification({
          title: 'Успешно!',
          message: `Группа ${value} добавлена`,
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
        fetchGroups();
      } catch (error) {
        Store.addNotification({
          title: 'Ошибка!',
          message: `Ошибка при добавлении группы ${value}`,
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
        <h4 className="card-title">Сообщества</h4>
        <h6 className="card-subtitle">на мониторинге {groups ? groups.length : 0}</h6>
        <InputField placeholder={'https://vk.com/club1'} onClick={addGroup} />
        {groups && groups.length > 0
          ? groups.map((group) => (
              <GroupElement
                apiUrl={apiUrl}
                key={group.name}
                group={group}
                onDeleteLocally={deleteGroupLocally}
              />
            ))
          : (groups === null || groups.length === 0) && <p>Нет добавленных групп</p>}
      </div>
    </div>
  );
}
