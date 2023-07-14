import React, { useState, useEffect } from 'react';
import axios from 'axios';
import PostElement from './PostElement';
import { VscRefresh } from 'react-icons/vsc';
import { Store } from 'react-notifications-component';

export default function PostsList({ apiUrl }) {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const timer = setInterval(() => {
      fetchPosts();
    }, 5000);

    fetchPosts();

    // Остановить таймер при размонтировании компонента
    return () => {
      clearInterval(timer);
    };
  }, []);

  const fetchPosts = async () => {
    try {
      const response = await axios.get(`${apiUrl}v1/posts/get?status=false`);
      setPosts(response.data);
    } catch (error) {
      console.error('Ошибка при получении списка постов:', error);
    }
  };

  const syncPosts = async () => {
    try {
      const r = await axios.post(`${apiUrl}v1/posts/sync`);
      fetchPosts();
      if (r.data.length == 0) {
        Store.addNotification({
          title: 'Успешно!',
          message: `Нет новых постов`,
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
      } else {
        Store.addNotification({
          title: 'Успешно!',
          message: `Загружено ${r.data.length} постов`,
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
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="card">
      <div className="card-body">
        <h4 className="card-title">
          Посты{' '}
          <span className="sync" onClick={syncPosts}>
            <VscRefresh />
          </span>
        </h4>
        <h6 className="card-subtitle">не прочитано {posts ? posts.length : 0}</h6>
        {posts && posts.length > 0
          ? posts.map((post) => (
              <PostElement apiUrl={apiUrl} key={post.hash} post={post} fetchPosts={fetchPosts} />
            ))
          : (posts === null || posts.length === 0) && <p>Нет постов</p>}
      </div>
    </div>
  );
}
