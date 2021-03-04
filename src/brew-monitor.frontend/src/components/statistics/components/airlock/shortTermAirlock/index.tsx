import React, { FC } from 'react';
import LineGraph from '../../../../common/graphs/line';
import Box from '@material-ui/core/Box';
import Respirator from '../../respirator';
import Airlock from '../../../../../types/Airlock';
import './styles.scss';

type ShortTermAirlockProps = {
  airlock: Airlock;
};

const ShortTermAirlockGraph: FC = () => {
  return (
    <Box className="shortTermAirlockGraph">
      <Respirator frequency={120} />
      <LineGraph />
    </Box>
  );
};

export default ShortTermAirlockGraph;
