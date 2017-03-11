%This script servers as a counter solution to the MachineLearning golang
%package mlearn
X = [1,  16663.2,   2.2704e+13,  316.205,  3454.65;
    1,   7308.8,  5.24519e+12,  261.312,  1461.75;
    1 ,  4040.7,  2.33762e+12,  235.164,   851.81;
    1,     6174 , 4.46038e+12, 251.802,  1324.23;
    1,  11510.7,  9.22067e+12,   289.79,   2159.9;
    1,   5979.6, 4.30443e+12,   248.71,  1252.99;
    1,   2632.1,  1.99276e+12,  224.107,   504.03;
    1,   6539.3 , 4.80281e+12,  254.933,  1381.53;
    1,     3345,   2.2705e+12,  230.815,   745.74;
    1,   9089.2,  6.70789e+12,  274.552,  1652.46;
    1,  14718.6,  1.94421e+13,  303.492,  2982.54;
    1,   7664.1,  6.23143e+12,  264.561,  1515.74;
    1,  14964.4,  1.83128e+13,  309.348,  3457.08;
    1,  17348.1,  2.31958e+13,  318.563,  3506.11;
    1,  14418.7,  1.53953e+13,  306.202,  3517.68;
    1,   5657.7,  3.73523e+12,  246.399,  1143.74;
    1,    17947,  2.06127e+13,  320.897,  3688.29;
    1,   4590.2,  2.61331e+12,  239.595,   990.38;
    1,   5252.6,  3.45802e+12,   244.11,  1064.42;
    1,  14477.6,  1.68302e+13,  300.807,  2728.69;
    1,     3211,  2.41659e+12,   228.67,   678.24;
    1,  13855.9,  1.45476e+13,  298.145,  2655.05;
    1,  10621.8,   7.6292e+12,  284.184,  1862.85;
    1,  16155.3,  2.21276e+13,  313.998,  3536.95;
    1,   6878.7,  4.72283e+12,  258.103,  1409.39;
    1,  12274.9,  1.12092e+13,  292.635,  2292.84;
    1,     2086,  1.36905e+12,  219.307,   409.22;
    1,  15517.9,  2.18233e+13,  311.663,  3603.06;
    1,  13093.7,  1.27193e+13,  295.507,  2471.96;
    1,  10977.5,  7.88783e+12,  286.974,  2010.89;
    1,  10284.8,  7.89202e+12,  281.422,  1788.95;
    1,   3638.1,  2.22501e+12,  232.979,   808.36;
    1,   2356.6,  1.58647e+12,  221.694,   458.75;
    1,   9660.6,  7.01544e+12,  277.966,  1701.84;
    1,   4870.2,   3.0576e+12,  241.842,  1004.02;
    1,   4346.7,  2.36575e+12,  237.369,   946.34;
    1,   2862.5,  2.40821e+12,  226.546,   590.94;
    1,   8608.5,   6.7342e+12 ,  271.18,  1601.12;
    1,   8100.2,   6.5748e+12 ,  267.85,  1560.48];

y = [  90.25;
    2028.18;
      285.4;
      110.9;
      117.3;
      250.5;
     416.08;
     895.84;
    1425.59;
    1378.76;
    1822.36;
    1140.21;
     465.25;
    1123.58;
     766.22;
     1480.4;
    1132.52;
     865.58;
     339.97;
     325.49;
      208.2;
    1278.73;
      264.5;
    1424.16;
      166.4;
    1282.62;
      99.71;
    1300.58;
    1335.63;
     614.42;
    1181.41;
      171.6;
     472.99;
      103.8;
        133;
    1248.77;
     435.23;
      144.3;
     963.36];
 
 sizeX = size(X);
 thetaRows = sizeX(2);
 
 theta = zeros(thetaRows, 1);
 
 CostHistory; theta = gradientDescent(X, theta, y, .01, 800);
 
 X = [1, 2, 3; 2, 5, 9; 3, 5, 10];
 y = [6; 16; 18];
 
 sizeX = size(X);
 thetaRows = sizeX(2);
 
 theta2 = zeros(thetaRows, 1);
 
 CostHistory2; theta2 = gradientDescent(X, theta2, y, .01, 6400);
 
 fprintf('THETA: \n');
 disp(theta);
 
 fprintf('\n\n');
 
 fprintf('THETA2: \n');
 disp(theta2);
