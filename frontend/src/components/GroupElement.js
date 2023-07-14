import React from 'react';
import axios from 'axios';

function GroupElement({ group, onDeleteLocally, apiUrl }) {
  const handleDeleteGroup = async () => {
    try {
      await axios.post(`${apiUrl}v1/group/delete`, { address: group.full_url });
      onDeleteLocally(group);
    } catch (error) {
      console.error('Произошла ошибка при удалении группы:', error);
    }
  };

  return (
    <div className="item-g pb-3 d-flex align-items-center comment-row ">
      <img
        className="btn-circle circle-image d-flex align-items-center"
        src={group.photo_url}
        alt=""
      />
      <div className="ms-3">
        <h5 className="mb-0 fw-bold">{group.name}</h5>
        <span className="text-muted fs-6">
          <a href={group.full_url}>{group.full_url}</a>
        </span>
        <br />
        <span className="action-icons ">
          <a href="#" onClick={handleDeleteGroup}>
            удалить
          </a>
        </span>
      </div>
    </div>
  );
}

export default GroupElement;
