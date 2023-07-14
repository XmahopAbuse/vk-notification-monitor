import * as React from 'react';
import GroupList from '../components/GroupsList';
import PostsList from '../components/PostsList';
import KeywordsList from '../components/KeywordsList';
import { VscSettingsGear } from 'react-icons/vsc';

export default function Home() {
  const apiUrl = process.env.REACT_APP_API_URL || '';
  return (
    <div>
      <div className="row">
        <div style={{ textAlign: 'right' }} className="col-lg-12">
          <a style={{ marginRight: '40px', marginTop: '20px', fontSize: '30px' }} href="/settings">
            <VscSettingsGear />
          </a>
        </div>
      </div>
      <div style={{ padding: '40px' }} className="page-wrapper">
        <div className="row">
          <div className="col-lg-5">
            <PostsList apiUrl={apiUrl} />
          </div>
          <div className="col-lg-3">
            <GroupList apiUrl={apiUrl} />
          </div>
          <div className="col-lg-3">
            <KeywordsList apiUrl={apiUrl} />
          </div>
        </div>
      </div>
    </div>
  );
}
