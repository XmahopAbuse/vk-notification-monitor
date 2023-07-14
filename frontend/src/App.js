import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import { ReactNotifications } from 'react-notifications-component';
import 'react-notifications-component/dist/theme.css';
import SettingsPage from './pages/Settings';

function App() {
  return (
    <Router>
      <ReactNotifications />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/settings" element={<SettingsPage />} />
      </Routes>
    </Router>
  );
}

export default App;
