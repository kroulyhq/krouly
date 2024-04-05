import { useState, useEffect } from 'react';

const useKroulyStorage = () => {
  const [storageData, setStorageData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('../../storage/cryptodata.json'); // Ruta del archivo JSON
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setStorageData(data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  return storageData;
};

export default useKroulyStorage;
