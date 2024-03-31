import { useState, useEffect } from 'react';

interface DataItem {
  symbol: string;
  price: string;
}

const useKroulyStorage = (): DataItem[] => {
  const [storageData, setStorageData] = useState<DataItem[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('../../storage/cryptodata.json')
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data: DataItem[] = await response.json()
        setStorageData(data)
      } catch (error) {
        console.error('Error fetching data:', error)
      }
    };

    fetchData()
  }, []);

  return storageData
};

export default useKroulyStorage
