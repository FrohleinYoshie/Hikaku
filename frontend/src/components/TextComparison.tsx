import React, { useState } from 'react';
import axios from 'axios';

const TextComparison: React.FC = () => {
  const [text1, setText1] = useState('');
  const [text2, setText2] = useState('');
  const [comparison, setComparison] = useState('');
  const [text1Html, setText1Html] = useState('');
  const [text2Html, setText2Html] = useState('');

  const handleCompare = async () => {
    try {
      const response = await axios.post('http://localhost:8080/compare', { text1, text2 });
      setComparison(response.data.comparison);
      setText1Html(response.data.text1Html);
      setText2Html(response.data.text2Html);
    } catch (error) {
      console.error('Error comparing texts:', error);
      setComparison('Error occurred while comparing texts');
      setText1Html('');
      setText2Html('');
    }
  };

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <textarea
          value={text1}
          onChange={(e) => setText1(e.target.value)}
          placeholder="文章を入れてください"
          style={{ width: '45%', height: '150px' }}
        />
        <textarea
          value={text2}
          onChange={(e) => setText2(e.target.value)}
          placeholder="文章を入れてください"
          style={{ width: '45%', height: '150px' }}
        />
      </div>
      <button onClick={handleCompare} style={{ margin: '10px 0' }}>比較する</button>
      <div style={{ display: 'flex', justifyContent: 'space-between' }}>
        <div style={{ width: '45%' }}>
          <h4>文章1:</h4>
          <div 
            dangerouslySetInnerHTML={{ __html: text1Html }} 
            style={{ whiteSpace: 'pre-wrap', fontFamily: 'monospace' }}
          />
        </div>
        <div style={{ width: '45%' }}>
          <h4>文章2:</h4>
          <div 
            dangerouslySetInnerHTML={{ __html: text2Html }} 
            style={{ whiteSpace: 'pre-wrap', fontFamily: 'monospace' }}
          />
        </div>
      </div>
    </div>
  );
};

export default TextComparison;