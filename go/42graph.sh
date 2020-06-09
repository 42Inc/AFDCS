#!/usr/bin/gnuplot
set terminal png size 1920,1080 enhanced font 'Arial, 16'

set style line 1 linecolor rgb 'red' linetype 1 linewidth 2
set style line 2 linecolor rgb 'blue' linetype 1 linewidth 2
set style line 3 linecolor rgb 'green' linetype 1 linewidth 2
set style line 4 linecolor rgb 'cyan' linetype 1 linewidth 2

set border linewidth 1
set key top left
set grid
set mytics
set mxtics
set format y "%.0f"
set xlabel "t" font "Arial, 16"
set format x "%.0f"
set ylabel "n" font "Arial, 16"
set xtics font "Arial, 12"
set ytics font "Arial, 12"
set rmargin 4
set tmargin 2

# set output 'data/afrvs_1_M.png'
# plot "data/afrvs_1_DT.dat" using 1:2 title "M +- qsrt(D)" with linespoints ls 3, \
#      "data/afrvs_1_DT.dat" using 1:3 notitle with linespoints ls 3,\
#      "data/afrvs_1_MT.dat" using 1:2 title "M" with linespoints ls 1, \
#      "data/afrvs_1_DP.dat" using 1:2 title "M +- qsrt(D) Experimental" with linespoints ls 4, \
#      "data/afrvs_1_DP.dat" using 1:3 notitle with linespoints ls 4, \
#      "data/afrvs_1_MP.dat" using 1:2 title "M Experimental" with linespoints ls 2

# set output 'data/afrvs_2_M.png'
# plot "data/afrvs_2_DT.dat" using 1:2 title "M +- qsrt(D)" with linespoints ls 3, \
#      "data/afrvs_2_DT.dat" using 1:3 notitle with linespoints ls 3,\
#      "data/afrvs_2_MT.dat" using 1:2 title "M" with linespoints ls 1, \
#      "data/afrvs_2_DP.dat" using 1:2 title "M +- qsrt(D) Experimental" with linespoints ls 4, \
#      "data/afrvs_2_DP.dat" using 1:3 notitle with linespoints ls 4, \
#      "data/afrvs_2_MP.dat" using 1:2 title "M Experimental" with linespoints ls 2

set output 'data/afrvs_3_M.png'
plot "data/afrvs_3_DT.dat" using 1:2 title "M +- qsrt(D)" with linespoints ls 3, \
     "data/afrvs_3_DT.dat" using 1:3 notitle with linespoints ls 3,\
     "data/afrvs_3_MT.dat" using 1:2 title "M" with linespoints ls 1, \
     "data/afrvs_3_DP.dat" using 1:2 title "M +- qsrt(D) Experimental" with linespoints ls 4, \
     "data/afrvs_3_DP.dat" using 1:3 notitle with linespoints ls 4, \
     "data/afrvs_3_MP.dat" using 1:2 title "M Experimental" with linespoints ls 2

# set output 'data/afrvs_4_M.png'
# plot "data/afrvs_4_DT.dat" using 1:2 title "M +- qsrt(D)" with linespoints ls 3, \
#      "data/afrvs_4_DT.dat" using 1:3 notitle with linespoints ls 3,\
#      "data/afrvs_4_MT.dat" using 1:2 title "M" with linespoints ls 1, \
#      "data/afrvs_4_DP.dat" using 1:2 title "M +- qsrt(D) Experimental" with linespoints ls 4, \
#      "data/afrvs_4_DP.dat" using 1:3 notitle with linespoints ls 4, \
#      "data/afrvs_4_MP.dat" using 1:2 title "M Experimental" with linespoints ls 2
