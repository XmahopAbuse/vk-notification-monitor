import React, { useEffect, useState } from 'react';
import axios from 'axios';
import DateComponent from './Date';
import { BiLinkExternal } from 'react-icons/bi';
import { MdDone } from 'react-icons/md';

function PostElement({ post, fetchPosts, apiUrl }) {
  const [authorData, setAuthorData] = useState(null);

  const getPostAuthor = async () => {
    try {
      const response = await axios.post(
        `${apiUrl}v1/author/get`,
        {
          author_id: post.from_id.toString(),
        },
        {
          timeout: 1000000,
        },
      );
      setAuthorData(response.data);
    } catch (error) {
      if (axios.isCancel(error)) {
        console.log('Запрос отменен:', error.message);
      } else {
        console.error('Произошла ошибка при получении пользователя', error);
      }
    }
  };

  const onClickClosePost = async (event) => {
    event.preventDefault();
    try {
      await axios.patch(`${apiUrl}v1/posts/update/${post.hash}`, {
        status: true,
      });
      fetchPosts();
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    getPostAuthor();
  }, [post.from_id]);

  return (
    <div className="d-flex flex-row comment-row m-t-0">
      <div className="p-2">
        {authorData && (
          <img
            src={authorData.Photo}
            alt="user"
            width="60"
            height="60"
            className="rounded-circle"
            style={{ marginTop: '15px' }}
          />
        )}
        {!authorData && (
          <img
            src="https://static.vecteezy.com/system/resources/thumbnails/003/337/584/small/default-avatar-photo-placeholder-profile-icon-vector.jpg"
            alt="user"
            width="60"
            height="60"
            className="rounded-circle"
          />
        )}
      </div>
      {authorData && (
        <div className="comment-text w-100">
          <h6 className="font-medium">
            <a href={authorData.FullUrl}>{authorData.Name}</a>
          </h6>
          <span className="m-b-15 d-block">{post.text} </span>
          <div className="comment-footer">
            <span className="text-muted float-end mt-3">
              <DateComponent date={post.date} />
            </span>
            <a
              target="_blank"
              title="Перейти на страницу с постом"
              href={`${post.post_url}`}
              className="mt-3 post-icon">
              <BiLinkExternal />
            </a>
            <a
              style={{ marginLeft: '10px' }}
              onClick={onClickClosePost}
              href="#"
              target="_blank"
              title="Отвечено"
              className="mt-3 post-icon">
              <MdDone />
            </a>
          </div>
        </div>
      )}
      {!authorData && (
        <div className="comment-text w-100">
          <h6 className="font-medium">
            <a href="#">Не удалось загрузить</a>
          </h6>
          <span className="m-b-15 d-block">{post.text} </span>
          <div className="comment-footer">
            <span className="text-muted float-end">
              <DateComponent date={post.date} />
            </span>
            <a
              target="_blank"
              title="Перейти на страницу с постом"
              href={`${post.post_url}`}
              className="mt-3 post-icon">
              <BiLinkExternal />
            </a>
            <a
              style={{ marginLeft: '10px' }}
              onClick={onClickClosePost}
              href="#"
              target="_blank"
              title="Отвечено"
              className="mt-3 post-icon">
              <MdDone />
            </a>
          </div>
        </div>
      )}
    </div>
  );
}

export default PostElement;
