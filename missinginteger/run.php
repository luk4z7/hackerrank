<?php 

// you can write to stdout for debugging purposes, e.g.
// print "this is a debug message\n";

function solution($A) {
    $a = $A;
    /**
     * Assumes that the first positive integer
     * in an empty array will be 1 - also stops
     * potential for infinite loop in while()
     * below
     */
    if(!count($a)){
        return 1;
    }

    /**
     * Lowest possible number in set
     * -1 so it starts counting at
     * -1000000
     */
    $i =-1000001;
    $c = 1;

    while($c){
        $i++;
        if(($i > 0 && !in_array($i, $a)) || $i > 1000000){
            $c = 0;
        }
    }

    /**
     * If $i is negative, the lowest positive
     * integer in set will be 1, otherwise it
     * will be the set value of $i unless $i
     * is out of range for this task
     */
    if($i <= 1000000){
        return ($i < 1) ? 1 : $i;
    }
    return 'Value out of range';
}


$data = [2, 3, 7, 6, 8, -1, -10, 15];
echo solution($data) . PHP_EOL;

$data = [2, 3, -7, 6, 8, 1, -10, 15];
echo solution($data) . PHP_EOL;

$data = [1, 1, 0, -1, -2];
echo solution($data) . PHP_EOL;


