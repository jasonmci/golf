import { Player } from './models/player';
import { Course } from './models/course';
import { rollDice } from './utils/dice';
import readlineSync = require('readline-sync');
import * as fs from 'fs';
import * as path from 'path';

const loadCourseData = (filePath: string): Course => {
    const fullPath = path.join(__dirname, filePath);
    const data = fs.readFileSync(fullPath, 'utf-8');
    const jsonData = JSON.parse(data);

    return new Course(jsonData.name, jsonData.holes);
}

const loadPlayerData = (filePath: string): Player => {
    const fullPath = path.join(__dirname, filePath);
    const data = fs.readFileSync(fullPath, 'utf-8');
    const jsonData = JSON.parse(data);

    return new Player(jsonData.name, jsonData.country, jsonData.qp, jsonData.rolls);
};

const main = () => {
    console.log("Welcome to TS Golf");

    const course = loadCourseData('../courses/pebble-beach.json');
    // get the player data next.
    console.log(`Playing on ${course.name}`);

    const tigerWoods = loadPlayerData('../player-cards/tiger-woods.json')
    tigerWoods.printPlayerInfo();
    
    // what hole number are you playing?
    const holeNumber = readlineSync.question('Enter hole number to view outcomes: ');
    course.printHoleOutcomes(holeNumber);

    const diceRoll = rollDice();
    console.log(`You rolled a ${diceRoll}`);
}

main();
