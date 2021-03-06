import { useState, useEffect } from 'react';
import Brew from '../types/Brew';

const header = {
  headers: {
    'Content-type': 'application/json',
  },
};

const useBrew = () => {
  const [brews, setBrews] = useState<Brew[]>([]);
  const [brew, setBrew] = useState<Brew>(brews[brews.length - 1]);

  useEffect(() => {
    fetchBrews();
  }, []);

  useEffect(() => {
    setBrew(brews[brews.length - 1]);
  }, [brews]);

  const fetchBrews = async (): Promise<void> => {
    try {
      const response = await fetch(`http://localhost:8080/brews`, header).then((response) => {
        if (!response.ok) {
          throw new Error(response.statusText);
        }
        return response;
      });
      const data = (await response.json()) as Brew[];
      setBrews(data);
    } catch (e) {
      console.log(`Error while fetching brews: ${e}`);
    }
  };

  const fetchBrew = async (brewId: string): Promise<void> => {
    try {
      const response = await fetch(`http://localhost:8080/brew/${brewId}`, header).then((response) => {
        if (!response.ok) {
          throw new Error(response.statusText);
        }
        return response;
      });
      const data = (await response.json()) as Brew;
      setBrew(data);
    } catch (e) {
      console.log(`Error while fetching brewId ${brewId}: ${e}`);
    }
  };

  return {
    brew,
    brews,
    fetchBrew,
    fetchBrews,
  };
};

export default useBrew;
