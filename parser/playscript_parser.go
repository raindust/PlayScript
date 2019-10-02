// Code generated from PlayScript.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // PlayScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 115, 537,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 3, 2,
	3, 2, 7, 2, 90, 10, 2, 12, 2, 14, 2, 93, 11, 2, 3, 3, 3, 3, 5, 3, 97, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 103, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 5, 6, 112, 10, 6, 3, 6, 3, 6, 7, 6, 116, 10, 6, 12, 6, 14,
	6, 119, 11, 6, 3, 7, 3, 7, 3, 7, 7, 7, 124, 10, 7, 12, 7, 14, 7, 127, 11,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 136, 10, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 152, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 5, 8, 190, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 202, 10, 8, 12, 8, 14, 8, 205, 11, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 210, 10, 9, 12, 9, 14, 9, 213, 11, 9, 3, 10, 3,
	10, 3, 10, 5, 10, 218, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 224,
	10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 230, 10, 10, 3, 10, 5, 10, 233,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11,
	243, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 255, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 267, 10, 16, 3, 17, 7, 17, 270, 10,
	17, 12, 17, 14, 17, 273, 11, 17, 3, 18, 3, 18, 5, 18, 277, 10, 18, 3, 19,
	3, 19, 3, 19, 7, 19, 282, 10, 19, 12, 19, 14, 19, 285, 11, 19, 3, 20, 3,
	20, 3, 20, 7, 20, 290, 10, 20, 12, 20, 14, 20, 293, 11, 20, 3, 21, 5, 21,
	296, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 302, 10, 21, 12, 21, 14,
	21, 305, 11, 21, 3, 21, 3, 21, 5, 21, 309, 10, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 5, 22, 315, 10, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 7, 23, 322,
	10, 23, 12, 23, 14, 23, 325, 11, 23, 3, 23, 3, 23, 5, 23, 329, 10, 23,
	3, 23, 5, 23, 332, 10, 23, 3, 24, 7, 24, 335, 10, 24, 12, 24, 14, 24, 338,
	11, 24, 3, 24, 3, 24, 3, 24, 3, 25, 7, 25, 344, 10, 25, 12, 25, 14, 25,
	347, 11, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 5, 27, 359, 10, 27, 3, 27, 3, 27, 5, 27, 363, 10, 27, 3, 27,
	3, 27, 3, 28, 3, 28, 7, 28, 369, 10, 28, 12, 28, 14, 28, 372, 11, 28, 3,
	28, 3, 28, 3, 29, 3, 29, 5, 29, 378, 10, 29, 3, 30, 3, 30, 3, 30, 5, 30,
	383, 10, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3,
	33, 5, 33, 394, 10, 33, 3, 33, 3, 33, 5, 33, 398, 10, 33, 3, 33, 3, 33,
	5, 33, 402, 10, 33, 3, 34, 3, 34, 5, 34, 406, 10, 34, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 7, 36, 416, 10, 36, 12, 36, 14,
	36, 419, 11, 36, 3, 36, 5, 36, 422, 10, 36, 5, 36, 424, 10, 36, 3, 36,
	3, 36, 3, 37, 3, 37, 5, 37, 430, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 7,
	38, 436, 10, 38, 12, 38, 14, 38, 439, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 40, 7, 40, 449, 10, 40, 12, 40, 14, 40, 452, 11,
	40, 3, 41, 6, 41, 455, 10, 41, 13, 41, 14, 41, 456, 3, 41, 6, 41, 460,
	10, 41, 13, 41, 14, 41, 461, 3, 42, 3, 42, 3, 42, 5, 42, 467, 10, 42, 3,
	42, 3, 42, 3, 42, 5, 42, 472, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 5, 43, 480, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 7, 43, 502, 10, 43, 12, 43, 14, 43, 505, 11, 43, 3,
	43, 7, 43, 508, 10, 43, 12, 43, 14, 43, 511, 11, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 5, 43, 517, 10, 43, 3, 43, 3, 43, 3, 43, 5, 43, 522, 10, 43, 3,
	43, 3, 43, 3, 43, 5, 43, 527, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 5, 43, 535, 10, 43, 3, 43, 2, 3, 14, 44, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 2,
	13, 11, 2, 5, 5, 7, 7, 10, 10, 16, 16, 22, 22, 29, 29, 31, 31, 39, 39,
	54, 54, 3, 2, 87, 90, 3, 2, 77, 78, 4, 2, 91, 92, 96, 96, 3, 2, 89, 90,
	4, 2, 75, 76, 82, 83, 4, 2, 81, 81, 84, 84, 4, 2, 74, 74, 97, 107, 3, 2,
	87, 88, 3, 2, 55, 58, 3, 2, 59, 60, 2, 592, 2, 86, 3, 2, 2, 2, 4, 96, 3,
	2, 2, 2, 6, 98, 3, 2, 2, 2, 8, 106, 3, 2, 2, 2, 10, 111, 3, 2, 2, 2, 12,
	120, 3, 2, 2, 2, 14, 135, 3, 2, 2, 2, 16, 206, 3, 2, 2, 2, 18, 232, 3,
	2, 2, 2, 20, 242, 3, 2, 2, 2, 22, 244, 3, 2, 2, 2, 24, 246, 3, 2, 2, 2,
	26, 254, 3, 2, 2, 2, 28, 256, 3, 2, 2, 2, 30, 266, 3, 2, 2, 2, 32, 271,
	3, 2, 2, 2, 34, 276, 3, 2, 2, 2, 36, 278, 3, 2, 2, 2, 38, 286, 3, 2, 2,
	2, 40, 295, 3, 2, 2, 2, 42, 312, 3, 2, 2, 2, 44, 331, 3, 2, 2, 2, 46, 336,
	3, 2, 2, 2, 48, 345, 3, 2, 2, 2, 50, 352, 3, 2, 2, 2, 52, 354, 3, 2, 2,
	2, 54, 366, 3, 2, 2, 2, 56, 377, 3, 2, 2, 2, 58, 382, 3, 2, 2, 2, 60, 384,
	3, 2, 2, 2, 62, 387, 3, 2, 2, 2, 64, 401, 3, 2, 2, 2, 66, 405, 3, 2, 2,
	2, 68, 407, 3, 2, 2, 2, 70, 411, 3, 2, 2, 2, 72, 429, 3, 2, 2, 2, 74, 431,
	3, 2, 2, 2, 76, 440, 3, 2, 2, 2, 78, 445, 3, 2, 2, 2, 80, 454, 3, 2, 2,
	2, 82, 471, 3, 2, 2, 2, 84, 534, 3, 2, 2, 2, 86, 91, 7, 115, 2, 2, 87,
	88, 7, 73, 2, 2, 88, 90, 7, 115, 2, 2, 89, 87, 3, 2, 2, 2, 90, 93, 3, 2,
	2, 2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 3, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 94, 97, 5, 10, 6, 2, 95, 97, 7, 50, 2, 2, 96, 94, 3, 2, 2,
	2, 96, 95, 3, 2, 2, 2, 97, 5, 3, 2, 2, 2, 98, 99, 7, 53, 2, 2, 99, 100,
	5, 4, 3, 2, 100, 102, 7, 65, 2, 2, 101, 103, 5, 12, 7, 2, 102, 101, 3,
	2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105, 7, 66, 2,
	2, 105, 7, 3, 2, 2, 2, 106, 107, 9, 2, 2, 2, 107, 9, 3, 2, 2, 2, 108, 112,
	5, 2, 2, 2, 109, 112, 5, 6, 4, 2, 110, 112, 5, 8, 5, 2, 111, 108, 3, 2,
	2, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 117, 3, 2, 2, 2,
	113, 114, 7, 69, 2, 2, 114, 116, 7, 70, 2, 2, 115, 113, 3, 2, 2, 2, 116,
	119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 11, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 125, 5, 10, 6, 2, 121, 122, 7, 72,
	2, 2, 122, 124, 5, 10, 6, 2, 123, 121, 3, 2, 2, 2, 124, 127, 3, 2, 2, 2,
	125, 123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 13, 3, 2, 2, 2, 127, 125,
	3, 2, 2, 2, 128, 129, 8, 8, 1, 2, 129, 136, 5, 20, 11, 2, 130, 136, 5,
	18, 10, 2, 131, 132, 9, 3, 2, 2, 132, 136, 5, 14, 8, 17, 133, 134, 9, 4,
	2, 2, 134, 136, 5, 14, 8, 16, 135, 128, 3, 2, 2, 2, 135, 130, 3, 2, 2,
	2, 135, 131, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 203, 3, 2, 2, 2, 137,
	138, 12, 15, 2, 2, 138, 139, 9, 5, 2, 2, 139, 202, 5, 14, 8, 16, 140, 141,
	12, 14, 2, 2, 141, 142, 9, 6, 2, 2, 142, 202, 5, 14, 8, 15, 143, 151, 12,
	13, 2, 2, 144, 145, 7, 76, 2, 2, 145, 152, 7, 76, 2, 2, 146, 147, 7, 75,
	2, 2, 147, 148, 7, 75, 2, 2, 148, 152, 7, 75, 2, 2, 149, 150, 7, 75, 2,
	2, 150, 152, 7, 75, 2, 2, 151, 144, 3, 2, 2, 2, 151, 146, 3, 2, 2, 2, 151,
	149, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 202, 5, 14, 8, 14, 154, 155,
	12, 12, 2, 2, 155, 156, 9, 7, 2, 2, 156, 202, 5, 14, 8, 13, 157, 158, 12,
	10, 2, 2, 158, 159, 9, 8, 2, 2, 159, 202, 5, 14, 8, 11, 160, 161, 12, 9,
	2, 2, 161, 162, 7, 93, 2, 2, 162, 202, 5, 14, 8, 10, 163, 164, 12, 8, 2,
	2, 164, 165, 7, 95, 2, 2, 165, 202, 5, 14, 8, 9, 166, 167, 12, 7, 2, 2,
	167, 168, 7, 94, 2, 2, 168, 202, 5, 14, 8, 8, 169, 170, 12, 6, 2, 2, 170,
	171, 7, 85, 2, 2, 171, 202, 5, 14, 8, 7, 172, 173, 12, 5, 2, 2, 173, 174,
	7, 86, 2, 2, 174, 202, 5, 14, 8, 6, 175, 176, 12, 4, 2, 2, 176, 177, 7,
	79, 2, 2, 177, 178, 5, 14, 8, 2, 178, 179, 7, 80, 2, 2, 179, 180, 5, 14,
	8, 5, 180, 202, 3, 2, 2, 2, 181, 182, 12, 3, 2, 2, 182, 183, 9, 9, 2, 2,
	183, 202, 5, 14, 8, 3, 184, 185, 12, 21, 2, 2, 185, 189, 7, 73, 2, 2, 186,
	190, 7, 115, 2, 2, 187, 190, 5, 18, 10, 2, 188, 190, 7, 45, 2, 2, 189,
	186, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189, 188, 3, 2, 2, 2, 190, 202,
	3, 2, 2, 2, 191, 192, 12, 20, 2, 2, 192, 193, 7, 69, 2, 2, 193, 194, 5,
	14, 8, 2, 194, 195, 7, 70, 2, 2, 195, 202, 3, 2, 2, 2, 196, 197, 12, 18,
	2, 2, 197, 202, 9, 10, 2, 2, 198, 199, 12, 11, 2, 2, 199, 200, 7, 28, 2,
	2, 200, 202, 5, 10, 6, 2, 201, 137, 3, 2, 2, 2, 201, 140, 3, 2, 2, 2, 201,
	143, 3, 2, 2, 2, 201, 154, 3, 2, 2, 2, 201, 157, 3, 2, 2, 2, 201, 160,
	3, 2, 2, 2, 201, 163, 3, 2, 2, 2, 201, 166, 3, 2, 2, 2, 201, 169, 3, 2,
	2, 2, 201, 172, 3, 2, 2, 2, 201, 175, 3, 2, 2, 2, 201, 181, 3, 2, 2, 2,
	201, 184, 3, 2, 2, 2, 201, 191, 3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 201,
	198, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204,
	3, 2, 2, 2, 204, 15, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 211, 5, 14,
	8, 2, 207, 208, 7, 72, 2, 2, 208, 210, 5, 14, 8, 2, 209, 207, 3, 2, 2,
	2, 210, 213, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212,
	17, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 214, 215, 7, 115, 2, 2, 215, 217,
	7, 65, 2, 2, 216, 218, 5, 16, 9, 2, 217, 216, 3, 2, 2, 2, 217, 218, 3,
	2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 233, 7, 66, 2, 2, 220, 221, 7, 45,
	2, 2, 221, 223, 7, 65, 2, 2, 222, 224, 5, 16, 9, 2, 223, 222, 3, 2, 2,
	2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 233, 7, 66, 2, 2, 226,
	227, 7, 42, 2, 2, 227, 229, 7, 65, 2, 2, 228, 230, 5, 16, 9, 2, 229, 228,
	3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 7, 66,
	2, 2, 232, 214, 3, 2, 2, 2, 232, 220, 3, 2, 2, 2, 232, 226, 3, 2, 2, 2,
	233, 19, 3, 2, 2, 2, 234, 235, 7, 65, 2, 2, 235, 236, 5, 14, 8, 2, 236,
	237, 7, 66, 2, 2, 237, 243, 3, 2, 2, 2, 238, 243, 7, 45, 2, 2, 239, 243,
	7, 42, 2, 2, 240, 243, 5, 26, 14, 2, 241, 243, 7, 115, 2, 2, 242, 234,
	3, 2, 2, 2, 242, 238, 3, 2, 2, 2, 242, 239, 3, 2, 2, 2, 242, 240, 3, 2,
	2, 2, 242, 241, 3, 2, 2, 2, 243, 21, 3, 2, 2, 2, 244, 245, 9, 11, 2, 2,
	245, 23, 3, 2, 2, 2, 246, 247, 9, 12, 2, 2, 247, 25, 3, 2, 2, 2, 248, 255,
	5, 22, 12, 2, 249, 255, 5, 24, 13, 2, 250, 255, 7, 62, 2, 2, 251, 255,
	7, 63, 2, 2, 252, 255, 7, 61, 2, 2, 253, 255, 7, 64, 2, 2, 254, 248, 3,
	2, 2, 2, 254, 249, 3, 2, 2, 2, 254, 250, 3, 2, 2, 2, 254, 251, 3, 2, 2,
	2, 254, 252, 3, 2, 2, 2, 254, 253, 3, 2, 2, 2, 255, 27, 3, 2, 2, 2, 256,
	257, 7, 67, 2, 2, 257, 258, 5, 32, 17, 2, 258, 259, 7, 68, 2, 2, 259, 29,
	3, 2, 2, 2, 260, 261, 5, 74, 38, 2, 261, 262, 7, 71, 2, 2, 262, 267, 3,
	2, 2, 2, 263, 267, 5, 84, 43, 2, 264, 267, 5, 40, 21, 2, 265, 267, 5, 52,
	27, 2, 266, 260, 3, 2, 2, 2, 266, 263, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2,
	266, 265, 3, 2, 2, 2, 267, 31, 3, 2, 2, 2, 268, 270, 5, 30, 16, 2, 269,
	268, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272,
	3, 2, 2, 2, 272, 33, 3, 2, 2, 2, 273, 271, 3, 2, 2, 2, 274, 277, 5, 28,
	15, 2, 275, 277, 7, 71, 2, 2, 276, 274, 3, 2, 2, 2, 276, 275, 3, 2, 2,
	2, 277, 35, 3, 2, 2, 2, 278, 283, 7, 115, 2, 2, 279, 280, 7, 73, 2, 2,
	280, 282, 7, 115, 2, 2, 281, 279, 3, 2, 2, 2, 282, 285, 3, 2, 2, 2, 283,
	281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 37, 3, 2, 2, 2, 285, 283, 3,
	2, 2, 2, 286, 291, 5, 36, 19, 2, 287, 288, 7, 72, 2, 2, 288, 290, 5, 36,
	19, 2, 289, 287, 3, 2, 2, 2, 290, 293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2,
	291, 292, 3, 2, 2, 2, 292, 39, 3, 2, 2, 2, 293, 291, 3, 2, 2, 2, 294, 296,
	5, 4, 3, 2, 295, 294, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 3, 2,
	2, 2, 297, 298, 7, 115, 2, 2, 298, 303, 5, 42, 22, 2, 299, 300, 7, 69,
	2, 2, 300, 302, 7, 70, 2, 2, 301, 299, 3, 2, 2, 2, 302, 305, 3, 2, 2, 2,
	303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 308, 3, 2, 2, 2, 305,
	303, 3, 2, 2, 2, 306, 307, 7, 47, 2, 2, 307, 309, 5, 38, 20, 2, 308, 306,
	3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 5, 34,
	18, 2, 311, 41, 3, 2, 2, 2, 312, 314, 7, 65, 2, 2, 313, 315, 5, 44, 23,
	2, 314, 313, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316,
	317, 7, 66, 2, 2, 317, 43, 3, 2, 2, 2, 318, 323, 5, 46, 24, 2, 319, 320,
	7, 72, 2, 2, 320, 322, 5, 46, 24, 2, 321, 319, 3, 2, 2, 2, 322, 325, 3,
	2, 2, 2, 323, 321, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 328, 3, 2, 2,
	2, 325, 323, 3, 2, 2, 2, 326, 327, 7, 72, 2, 2, 327, 329, 5, 48, 25, 2,
	328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 332, 3, 2, 2, 2, 330,
	332, 5, 48, 25, 2, 331, 318, 3, 2, 2, 2, 331, 330, 3, 2, 2, 2, 332, 45,
	3, 2, 2, 2, 333, 335, 5, 50, 26, 2, 334, 333, 3, 2, 2, 2, 335, 338, 3,
	2, 2, 2, 336, 334, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 339, 3, 2, 2,
	2, 338, 336, 3, 2, 2, 2, 339, 340, 5, 10, 6, 2, 340, 341, 5, 78, 40, 2,
	341, 47, 3, 2, 2, 2, 342, 344, 5, 50, 26, 2, 343, 342, 3, 2, 2, 2, 344,
	347, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 348,
	3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 349, 5, 10, 6, 2, 349, 350, 7, 111,
	2, 2, 350, 351, 5, 78, 40, 2, 351, 49, 3, 2, 2, 2, 352, 353, 7, 20, 2,
	2, 353, 51, 3, 2, 2, 2, 354, 355, 7, 11, 2, 2, 355, 358, 7, 115, 2, 2,
	356, 357, 7, 19, 2, 2, 357, 359, 5, 10, 6, 2, 358, 356, 3, 2, 2, 2, 358,
	359, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 361, 7, 26, 2, 2, 361, 363,
	5, 12, 7, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 3, 2,
	2, 2, 364, 365, 5, 54, 28, 2, 365, 53, 3, 2, 2, 2, 366, 370, 7, 67, 2,
	2, 367, 369, 5, 56, 29, 2, 368, 367, 3, 2, 2, 2, 369, 372, 3, 2, 2, 2,
	370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 373, 3, 2, 2, 2, 372,
	370, 3, 2, 2, 2, 373, 374, 7, 68, 2, 2, 374, 55, 3, 2, 2, 2, 375, 378,
	7, 71, 2, 2, 376, 378, 5, 58, 30, 2, 377, 375, 3, 2, 2, 2, 377, 376, 3,
	2, 2, 2, 378, 57, 3, 2, 2, 2, 379, 383, 5, 40, 21, 2, 380, 383, 5, 60,
	31, 2, 381, 383, 5, 52, 27, 2, 382, 379, 3, 2, 2, 2, 382, 380, 3, 2, 2,
	2, 382, 381, 3, 2, 2, 2, 383, 59, 3, 2, 2, 2, 384, 385, 5, 74, 38, 2, 385,
	386, 7, 71, 2, 2, 386, 61, 3, 2, 2, 2, 387, 388, 7, 65, 2, 2, 388, 389,
	5, 14, 8, 2, 389, 390, 7, 66, 2, 2, 390, 63, 3, 2, 2, 2, 391, 402, 5, 76,
	39, 2, 392, 394, 5, 66, 34, 2, 393, 392, 3, 2, 2, 2, 393, 394, 3, 2, 2,
	2, 394, 395, 3, 2, 2, 2, 395, 397, 7, 71, 2, 2, 396, 398, 5, 14, 8, 2,
	397, 396, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399,
	400, 7, 71, 2, 2, 400, 402, 5, 16, 9, 2, 401, 391, 3, 2, 2, 2, 401, 393,
	3, 2, 2, 2, 402, 65, 3, 2, 2, 2, 403, 406, 5, 74, 38, 2, 404, 406, 5, 16,
	9, 2, 405, 403, 3, 2, 2, 2, 405, 404, 3, 2, 2, 2, 406, 67, 3, 2, 2, 2,
	407, 408, 5, 78, 40, 2, 408, 409, 7, 74, 2, 2, 409, 410, 5, 72, 37, 2,
	410, 69, 3, 2, 2, 2, 411, 423, 7, 67, 2, 2, 412, 417, 5, 72, 37, 2, 413,
	414, 7, 72, 2, 2, 414, 416, 5, 72, 37, 2, 415, 413, 3, 2, 2, 2, 416, 419,
	3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 421, 3, 2,
	2, 2, 419, 417, 3, 2, 2, 2, 420, 422, 7, 72, 2, 2, 421, 420, 3, 2, 2, 2,
	421, 422, 3, 2, 2, 2, 422, 424, 3, 2, 2, 2, 423, 412, 3, 2, 2, 2, 423,
	424, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 426, 7, 68, 2, 2, 426, 71,
	3, 2, 2, 2, 427, 430, 5, 70, 36, 2, 428, 430, 5, 14, 8, 2, 429, 427, 3,
	2, 2, 2, 429, 428, 3, 2, 2, 2, 430, 73, 3, 2, 2, 2, 431, 432, 5, 10, 6,
	2, 432, 437, 5, 68, 35, 2, 433, 434, 7, 72, 2, 2, 434, 436, 5, 68, 35,
	2, 435, 433, 3, 2, 2, 2, 436, 439, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2, 437,
	438, 3, 2, 2, 2, 438, 75, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 440, 441, 5,
	10, 6, 2, 441, 442, 5, 78, 40, 2, 442, 443, 7, 80, 2, 2, 443, 444, 5, 14,
	8, 2, 444, 77, 3, 2, 2, 2, 445, 450, 7, 115, 2, 2, 446, 447, 7, 69, 2,
	2, 447, 449, 7, 70, 2, 2, 448, 446, 3, 2, 2, 2, 449, 452, 3, 2, 2, 2, 450,
	448, 3, 2, 2, 2, 450, 451, 3, 2, 2, 2, 451, 79, 3, 2, 2, 2, 452, 450, 3,
	2, 2, 2, 453, 455, 5, 82, 42, 2, 454, 453, 3, 2, 2, 2, 455, 456, 3, 2,
	2, 2, 456, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 459, 3, 2, 2, 2,
	458, 460, 5, 30, 16, 2, 459, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461,
	459, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 81, 3, 2, 2, 2, 463, 466, 7,
	8, 2, 2, 464, 467, 5, 14, 8, 2, 465, 467, 7, 115, 2, 2, 466, 464, 3, 2,
	2, 2, 466, 465, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 472, 7, 80, 2, 2,
	469, 470, 7, 14, 2, 2, 470, 472, 7, 80, 2, 2, 471, 463, 3, 2, 2, 2, 471,
	469, 3, 2, 2, 2, 472, 83, 3, 2, 2, 2, 473, 535, 5, 28, 15, 2, 474, 475,
	7, 24, 2, 2, 475, 476, 5, 62, 32, 2, 476, 479, 5, 84, 43, 2, 477, 478,
	7, 17, 2, 2, 478, 480, 5, 84, 43, 2, 479, 477, 3, 2, 2, 2, 479, 480, 3,
	2, 2, 2, 480, 535, 3, 2, 2, 2, 481, 482, 7, 23, 2, 2, 482, 483, 7, 65,
	2, 2, 483, 484, 5, 64, 33, 2, 484, 485, 7, 66, 2, 2, 485, 486, 5, 84, 43,
	2, 486, 535, 3, 2, 2, 2, 487, 488, 7, 52, 2, 2, 488, 489, 5, 62, 32, 2,
	489, 490, 5, 84, 43, 2, 490, 535, 3, 2, 2, 2, 491, 492, 7, 15, 2, 2, 492,
	493, 5, 84, 43, 2, 493, 494, 7, 52, 2, 2, 494, 495, 5, 62, 32, 2, 495,
	496, 7, 71, 2, 2, 496, 535, 3, 2, 2, 2, 497, 498, 7, 43, 2, 2, 498, 499,
	5, 62, 32, 2, 499, 503, 7, 67, 2, 2, 500, 502, 5, 80, 41, 2, 501, 500,
	3, 2, 2, 2, 502, 505, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503, 504, 3, 2,
	2, 2, 504, 509, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 506, 508, 5, 82, 42,
	2, 507, 506, 3, 2, 2, 2, 508, 511, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509,
	510, 3, 2, 2, 2, 510, 512, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 512, 513,
	7, 68, 2, 2, 513, 535, 3, 2, 2, 2, 514, 516, 7, 38, 2, 2, 515, 517, 5,
	14, 8, 2, 516, 515, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 518, 3, 2, 2,
	2, 518, 535, 7, 71, 2, 2, 519, 521, 7, 6, 2, 2, 520, 522, 7, 115, 2, 2,
	521, 520, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 523, 3, 2, 2, 2, 523,
	535, 7, 71, 2, 2, 524, 526, 7, 13, 2, 2, 525, 527, 7, 115, 2, 2, 526, 525,
	3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 535, 7, 71,
	2, 2, 529, 535, 7, 71, 2, 2, 530, 535, 5, 14, 8, 2, 531, 532, 7, 115, 2,
	2, 532, 533, 7, 80, 2, 2, 533, 535, 5, 84, 43, 2, 534, 473, 3, 2, 2, 2,
	534, 474, 3, 2, 2, 2, 534, 481, 3, 2, 2, 2, 534, 487, 3, 2, 2, 2, 534,
	491, 3, 2, 2, 2, 534, 497, 3, 2, 2, 2, 534, 514, 3, 2, 2, 2, 534, 519,
	3, 2, 2, 2, 534, 524, 3, 2, 2, 2, 534, 529, 3, 2, 2, 2, 534, 530, 3, 2,
	2, 2, 534, 531, 3, 2, 2, 2, 535, 85, 3, 2, 2, 2, 60, 91, 96, 102, 111,
	117, 125, 135, 151, 189, 201, 203, 211, 217, 223, 229, 232, 242, 254, 266,
	271, 276, 283, 291, 295, 303, 308, 314, 323, 328, 331, 336, 345, 358, 362,
	370, 377, 382, 393, 397, 401, 405, 417, 421, 423, 429, 437, 450, 456, 461,
	466, 471, 479, 503, 509, 516, 521, 526, 534,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'abstract'", "'assert'", "'boolean'", "'break'", "'byte'", "'case'",
	"'catch'", "'char'", "'class'", "'const'", "'continue'", "'default'", "'do'",
	"'double'", "'else'", "'enum'", "'extends'", "'final'", "'finally'", "'float'",
	"'for'", "'if'", "'goto'", "'implements'", "'import'", "'instanceof'",
	"'int'", "'interface'", "'long'", "'native'", "'new'", "'package'", "'private'",
	"'protected'", "'public'", "'return'", "'short'", "'static'", "'strictfp'",
	"'super'", "'switch'", "'synchronized'", "'this'", "'throw'", "'throws'",
	"'transient'", "'try'", "'void'", "'volatile'", "'while'", "'function'",
	"'string'", "", "", "", "", "", "", "", "", "", "'null'", "'('", "')'",
	"'{'", "'}'", "'['", "']'", "';'", "','", "'.'", "'='", "'>'", "'<'", "'!'",
	"'~'", "'?'", "':'", "'=='", "'<='", "'>='", "'!='", "'&&'", "'||'", "'++'",
	"'--'", "'+'", "'-'", "'*'", "'/'", "'&'", "'|'", "'^'", "'%'", "'+='",
	"'-='", "'*='", "'/='", "'&='", "'|='", "'^='", "'%='", "'<<='", "'>>='",
	"'>>>='", "'->'", "'::'", "'@'", "'...'",
}
var symbolicNames = []string{
	"", "ABSTRACT", "ASSERT", "BOOLEAN", "BREAK", "BYTE", "CASE", "CATCH",
	"CHAR", "CLASS", "CONST", "CONTINUE", "DEFAULT", "DO", "DOUBLE", "ELSE",
	"ENUM", "EXTENDS", "FINAL", "FINALLY", "FLOAT", "FOR", "IF", "GOTO", "IMPLEMENTS",
	"IMPORT", "INSTANCEOF", "INT", "INTERFACE", "LONG", "NATIVE", "NEW", "PACKAGE",
	"PRIVATE", "PROTECTED", "PUBLIC", "RETURN", "SHORT", "STATIC", "STRICTFP",
	"SUPER", "SWITCH", "SYNCHRONIZED", "THIS", "THROW", "THROWS", "TRANSIENT",
	"TRY", "VOID", "VOLATILE", "WHILE", "FUNCTION", "STRING", "DECIMAL_LITERAL",
	"HEX_LITERAL", "OCT_LITERAL", "BINARY_LITERAL", "FLOAT_LITERAL", "HEX_FLOAT_LITERAL",
	"BOOL_LITERAL", "CHAR_LITERAL", "STRING_LITERAL", "NULL_LITERAL", "LPAREN",
	"RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI", "COMMA", "DOT",
	"ASSIGN", "GT", "LT", "BANG", "TILDE", "QUESTION", "COLON", "EQUAL", "LE",
	"GE", "NOTEQUAL", "AND", "OR", "INC", "DEC", "ADD", "SUB", "MUL", "DIV",
	"BITAND", "BITOR", "CARET", "MOD", "ADD_ASSIGN", "SUB_ASSIGN", "MUL_ASSIGN",
	"DIV_ASSIGN", "AND_ASSIGN", "OR_ASSIGN", "XOR_ASSIGN", "MOD_ASSIGN", "LSHIFT_ASSIGN",
	"RSHIFT_ASSIGN", "URSHIFT_ASSIGN", "ARROW", "COLONCOLON", "AT", "ELLIPSIS",
	"WS", "COMMENT", "LINE_COMMENT", "IDENTIFIER",
}

var ruleNames = []string{
	"classOrInterfaceType", "typeTypeOrVoid", "functionType", "primitiveType",
	"typeType", "typeList", "expression", "expressionList", "functionCall",
	"primary", "intLiteral", "floatLiteral", "literal", "block", "blockStatement",
	"blockStatements", "functionBody", "qualifiedName", "qualifiedNameList",
	"functionDeclaration", "formalParameters", "formalParameterList", "formalParameter",
	"lastFormalParameter", "variableModifier", "classDeclaration", "classBody",
	"classBodyDeclaration", "memberDecalaration", "fieldDeclaration", "parExpression",
	"forControl", "forInit", "variableDeclarator", "arrayInitializer", "variableInitializer",
	"variableDeclarators", "enhanceForControl", "variableDeclaratorId", "switchBlockStatementGroup",
	"switchLabel", "statement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PlayScriptParser struct {
	*antlr.BaseParser
}

func NewPlayScriptParser(input antlr.TokenStream) *PlayScriptParser {
	this := new(PlayScriptParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PlayScript.g4"

	return this
}

// PlayScriptParser tokens.
const (
	PlayScriptParserEOF               = antlr.TokenEOF
	PlayScriptParserABSTRACT          = 1
	PlayScriptParserASSERT            = 2
	PlayScriptParserBOOLEAN           = 3
	PlayScriptParserBREAK             = 4
	PlayScriptParserBYTE              = 5
	PlayScriptParserCASE              = 6
	PlayScriptParserCATCH             = 7
	PlayScriptParserCHAR              = 8
	PlayScriptParserCLASS             = 9
	PlayScriptParserCONST             = 10
	PlayScriptParserCONTINUE          = 11
	PlayScriptParserDEFAULT           = 12
	PlayScriptParserDO                = 13
	PlayScriptParserDOUBLE            = 14
	PlayScriptParserELSE              = 15
	PlayScriptParserENUM              = 16
	PlayScriptParserEXTENDS           = 17
	PlayScriptParserFINAL             = 18
	PlayScriptParserFINALLY           = 19
	PlayScriptParserFLOAT             = 20
	PlayScriptParserFOR               = 21
	PlayScriptParserIF                = 22
	PlayScriptParserGOTO              = 23
	PlayScriptParserIMPLEMENTS        = 24
	PlayScriptParserIMPORT            = 25
	PlayScriptParserINSTANCEOF        = 26
	PlayScriptParserINT               = 27
	PlayScriptParserINTERFACE         = 28
	PlayScriptParserLONG              = 29
	PlayScriptParserNATIVE            = 30
	PlayScriptParserNEW               = 31
	PlayScriptParserPACKAGE           = 32
	PlayScriptParserPRIVATE           = 33
	PlayScriptParserPROTECTED         = 34
	PlayScriptParserPUBLIC            = 35
	PlayScriptParserRETURN            = 36
	PlayScriptParserSHORT             = 37
	PlayScriptParserSTATIC            = 38
	PlayScriptParserSTRICTFP          = 39
	PlayScriptParserSUPER             = 40
	PlayScriptParserSWITCH            = 41
	PlayScriptParserSYNCHRONIZED      = 42
	PlayScriptParserTHIS              = 43
	PlayScriptParserTHROW             = 44
	PlayScriptParserTHROWS            = 45
	PlayScriptParserTRANSIENT         = 46
	PlayScriptParserTRY               = 47
	PlayScriptParserVOID              = 48
	PlayScriptParserVOLATILE          = 49
	PlayScriptParserWHILE             = 50
	PlayScriptParserFUNCTION          = 51
	PlayScriptParserSTRING            = 52
	PlayScriptParserDECIMAL_LITERAL   = 53
	PlayScriptParserHEX_LITERAL       = 54
	PlayScriptParserOCT_LITERAL       = 55
	PlayScriptParserBINARY_LITERAL    = 56
	PlayScriptParserFLOAT_LITERAL     = 57
	PlayScriptParserHEX_FLOAT_LITERAL = 58
	PlayScriptParserBOOL_LITERAL      = 59
	PlayScriptParserCHAR_LITERAL      = 60
	PlayScriptParserSTRING_LITERAL    = 61
	PlayScriptParserNULL_LITERAL      = 62
	PlayScriptParserLPAREN            = 63
	PlayScriptParserRPAREN            = 64
	PlayScriptParserLBRACE            = 65
	PlayScriptParserRBRACE            = 66
	PlayScriptParserLBRACK            = 67
	PlayScriptParserRBRACK            = 68
	PlayScriptParserSEMI              = 69
	PlayScriptParserCOMMA             = 70
	PlayScriptParserDOT               = 71
	PlayScriptParserASSIGN            = 72
	PlayScriptParserGT                = 73
	PlayScriptParserLT                = 74
	PlayScriptParserBANG              = 75
	PlayScriptParserTILDE             = 76
	PlayScriptParserQUESTION          = 77
	PlayScriptParserCOLON             = 78
	PlayScriptParserEQUAL             = 79
	PlayScriptParserLE                = 80
	PlayScriptParserGE                = 81
	PlayScriptParserNOTEQUAL          = 82
	PlayScriptParserAND               = 83
	PlayScriptParserOR                = 84
	PlayScriptParserINC               = 85
	PlayScriptParserDEC               = 86
	PlayScriptParserADD               = 87
	PlayScriptParserSUB               = 88
	PlayScriptParserMUL               = 89
	PlayScriptParserDIV               = 90
	PlayScriptParserBITAND            = 91
	PlayScriptParserBITOR             = 92
	PlayScriptParserCARET             = 93
	PlayScriptParserMOD               = 94
	PlayScriptParserADD_ASSIGN        = 95
	PlayScriptParserSUB_ASSIGN        = 96
	PlayScriptParserMUL_ASSIGN        = 97
	PlayScriptParserDIV_ASSIGN        = 98
	PlayScriptParserAND_ASSIGN        = 99
	PlayScriptParserOR_ASSIGN         = 100
	PlayScriptParserXOR_ASSIGN        = 101
	PlayScriptParserMOD_ASSIGN        = 102
	PlayScriptParserLSHIFT_ASSIGN     = 103
	PlayScriptParserRSHIFT_ASSIGN     = 104
	PlayScriptParserURSHIFT_ASSIGN    = 105
	PlayScriptParserARROW             = 106
	PlayScriptParserCOLONCOLON        = 107
	PlayScriptParserAT                = 108
	PlayScriptParserELLIPSIS          = 109
	PlayScriptParserWS                = 110
	PlayScriptParserCOMMENT           = 111
	PlayScriptParserLINE_COMMENT      = 112
	PlayScriptParserIDENTIFIER        = 113
)

// PlayScriptParser rules.
const (
	PlayScriptParserRULE_classOrInterfaceType      = 0
	PlayScriptParserRULE_typeTypeOrVoid            = 1
	PlayScriptParserRULE_functionType              = 2
	PlayScriptParserRULE_primitiveType             = 3
	PlayScriptParserRULE_typeType                  = 4
	PlayScriptParserRULE_typeList                  = 5
	PlayScriptParserRULE_expression                = 6
	PlayScriptParserRULE_expressionList            = 7
	PlayScriptParserRULE_functionCall              = 8
	PlayScriptParserRULE_primary                   = 9
	PlayScriptParserRULE_intLiteral                = 10
	PlayScriptParserRULE_floatLiteral              = 11
	PlayScriptParserRULE_literal                   = 12
	PlayScriptParserRULE_block                     = 13
	PlayScriptParserRULE_blockStatement            = 14
	PlayScriptParserRULE_blockStatements           = 15
	PlayScriptParserRULE_functionBody              = 16
	PlayScriptParserRULE_qualifiedName             = 17
	PlayScriptParserRULE_qualifiedNameList         = 18
	PlayScriptParserRULE_functionDeclaration       = 19
	PlayScriptParserRULE_formalParameters          = 20
	PlayScriptParserRULE_formalParameterList       = 21
	PlayScriptParserRULE_formalParameter           = 22
	PlayScriptParserRULE_lastFormalParameter       = 23
	PlayScriptParserRULE_variableModifier          = 24
	PlayScriptParserRULE_classDeclaration          = 25
	PlayScriptParserRULE_classBody                 = 26
	PlayScriptParserRULE_classBodyDeclaration      = 27
	PlayScriptParserRULE_memberDecalaration        = 28
	PlayScriptParserRULE_fieldDeclaration          = 29
	PlayScriptParserRULE_parExpression             = 30
	PlayScriptParserRULE_forControl                = 31
	PlayScriptParserRULE_forInit                   = 32
	PlayScriptParserRULE_variableDeclarator        = 33
	PlayScriptParserRULE_arrayInitializer          = 34
	PlayScriptParserRULE_variableInitializer       = 35
	PlayScriptParserRULE_variableDeclarators       = 36
	PlayScriptParserRULE_enhanceForControl         = 37
	PlayScriptParserRULE_variableDeclaratorId      = 38
	PlayScriptParserRULE_switchBlockStatementGroup = 39
	PlayScriptParserRULE_switchLabel               = 40
	PlayScriptParserRULE_statement                 = 41
)

// IClassOrInterfaceTypeContext is an interface to support dynamic dispatch.
type IClassOrInterfaceTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassOrInterfaceTypeContext differentiates from other interfaces.
	IsClassOrInterfaceTypeContext()
}

type ClassOrInterfaceTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassOrInterfaceTypeContext() *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classOrInterfaceType
	return p
}

func (*ClassOrInterfaceTypeContext) IsClassOrInterfaceTypeContext() {}

func NewClassOrInterfaceTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassOrInterfaceTypeContext {
	var p = new(ClassOrInterfaceTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classOrInterfaceType

	return p
}

func (s *ClassOrInterfaceTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassOrInterfaceTypeContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserIDENTIFIER)
}

func (s *ClassOrInterfaceTypeContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, i)
}

func (s *ClassOrInterfaceTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserDOT)
}

func (s *ClassOrInterfaceTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, i)
}

func (s *ClassOrInterfaceTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassOrInterfaceTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassOrInterfaceTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassOrInterfaceType(s)
	}
}

func (s *ClassOrInterfaceTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassOrInterfaceType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassOrInterfaceType() (localctx IClassOrInterfaceTypeContext) {
	localctx = NewClassOrInterfaceTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PlayScriptParserRULE_classOrInterfaceType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(85)
				p.Match(PlayScriptParserDOT)
			}
			{
				p.SetState(86)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeTypeOrVoidContext is an interface to support dynamic dispatch.
type ITypeTypeOrVoidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeOrVoidContext differentiates from other interfaces.
	IsTypeTypeOrVoidContext()
}

type TypeTypeOrVoidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeOrVoidContext() *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeTypeOrVoid
	return p
}

func (*TypeTypeOrVoidContext) IsTypeTypeOrVoidContext() {}

func NewTypeTypeOrVoidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeOrVoidContext {
	var p = new(TypeTypeOrVoidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeTypeOrVoid

	return p
}

func (s *TypeTypeOrVoidContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeOrVoidContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeTypeOrVoidContext) VOID() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserVOID, 0)
}

func (s *TypeTypeOrVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeOrVoidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeOrVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeTypeOrVoid(s)
	}
}

func (s *TypeTypeOrVoidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeTypeOrVoid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeTypeOrVoid() (localctx ITypeTypeOrVoidContext) {
	localctx = NewTypeTypeOrVoidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PlayScriptParserRULE_typeTypeOrVoid)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserFUNCTION, PlayScriptParserSTRING, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.TypeType()
		}

	case PlayScriptParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(PlayScriptParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionTypeContext is an interface to support dynamic dispatch.
type IFunctionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionTypeContext differentiates from other interfaces.
	IsFunctionTypeContext()
}

type FunctionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionTypeContext() *FunctionTypeContext {
	var p = new(FunctionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionType
	return p
}

func (*FunctionTypeContext) IsFunctionTypeContext() {}

func NewFunctionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionTypeContext {
	var p = new(FunctionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionType

	return p
}

func (s *FunctionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionTypeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFUNCTION, 0)
}

func (s *FunctionTypeContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FunctionTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FunctionTypeContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *FunctionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionType(s)
	}
}

func (s *FunctionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionType(s)
	}
}

func (s *FunctionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionType() (localctx IFunctionTypeContext) {
	localctx = NewFunctionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PlayScriptParserRULE_functionType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(PlayScriptParserFUNCTION)
	}
	{
		p.SetState(97)
		p.TypeTypeOrVoid()
	}
	{
		p.SetState(98)
		p.Match(PlayScriptParserLPAREN)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(99)
			p.TypeList()
		}

	}
	{
		p.SetState(102)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBOOLEAN, 0)
}

func (s *PrimitiveTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCHAR, 0)
}

func (s *PrimitiveTypeContext) BYTE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBYTE, 0)
}

func (s *PrimitiveTypeContext) SHORT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSHORT, 0)
}

func (s *PrimitiveTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINT, 0)
}

func (s *PrimitiveTypeContext) LONG() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLONG, 0)
}

func (s *PrimitiveTypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFLOAT, 0)
}

func (s *PrimitiveTypeContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOUBLE, 0)
}

func (s *PrimitiveTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSTRING, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PlayScriptParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || _la == PlayScriptParserSHORT || _la == PlayScriptParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeTypeContext is an interface to support dynamic dispatch.
type ITypeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeTypeContext differentiates from other interfaces.
	IsTypeTypeContext()
}

type TypeTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeTypeContext() *TypeTypeContext {
	var p = new(TypeTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeType
	return p
}

func (*TypeTypeContext) IsTypeTypeContext() {}

func NewTypeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeTypeContext {
	var p = new(TypeTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeType

	return p
}

func (s *TypeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeTypeContext) ClassOrInterfaceType() IClassOrInterfaceTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassOrInterfaceTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassOrInterfaceTypeContext)
}

func (s *TypeTypeContext) FunctionType() IFunctionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionTypeContext)
}

func (s *TypeTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *TypeTypeContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *TypeTypeContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *TypeTypeContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *TypeTypeContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *TypeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeType(s)
	}
}

func (s *TypeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeType(s)
	}
}

func (s *TypeTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeType() (localctx ITypeTypeContext) {
	localctx = NewTypeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PlayScriptParserRULE_typeType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserIDENTIFIER:
		{
			p.SetState(106)
			p.ClassOrInterfaceType()
		}

	case PlayScriptParserFUNCTION:
		{
			p.SetState(107)
			p.FunctionType()
		}

	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserSTRING:
		{
			p.SetState(108)
			p.PrimitiveType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(111)
				p.Match(PlayScriptParserLBRACK)
			}
			{
				p.SetState(112)
				p.Match(PlayScriptParserRBRACK)
			}

		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeListContext is an interface to support dynamic dispatch.
type ITypeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeListContext differentiates from other interfaces.
	IsTypeListContext()
}

type TypeListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeListContext() *TypeListContext {
	var p = new(TypeListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_typeList
	return p
}

func (*TypeListContext) IsTypeListContext() {}

func NewTypeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeListContext {
	var p = new(TypeListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_typeList

	return p
}

func (s *TypeListContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeListContext) AllTypeType() []ITypeTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem())
	var tst = make([]ITypeTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeTypeContext)
		}
	}

	return tst
}

func (s *TypeListContext) TypeType(i int) ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *TypeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *TypeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *TypeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterTypeList(s)
	}
}

func (s *TypeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitTypeList(s)
	}
}

func (s *TypeListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitTypeList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) TypeList() (localctx ITypeListContext) {
	localctx = NewTypeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PlayScriptParserRULE_typeList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.TypeType()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(119)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(120)
			p.TypeType()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetBop returns the bop token.
	GetBop() antlr.Token

	// GetPostfix returns the postfix token.
	GetPostfix() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetBop sets the bop token.
	SetBop(antlr.Token)

	// SetPostfix sets the postfix token.
	SetPostfix(antlr.Token)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	prefix  antlr.Token
	bop     antlr.Token
	postfix antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetPrefix() antlr.Token { return s.prefix }

func (s *ExpressionContext) GetBop() antlr.Token { return s.bop }

func (s *ExpressionContext) GetPostfix() antlr.Token { return s.postfix }

func (s *ExpressionContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *ExpressionContext) SetBop(v antlr.Token) { s.bop = v }

func (s *ExpressionContext) SetPostfix(v antlr.Token) { s.postfix = v }

func (s *ExpressionContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUB, 0)
}

func (s *ExpressionContext) INC() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINC, 0)
}

func (s *ExpressionContext) DEC() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDEC, 0)
}

func (s *ExpressionContext) TILDE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTILDE, 0)
}

func (s *ExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBANG, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMOD, 0)
}

func (s *ExpressionContext) AllLT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLT)
}

func (s *ExpressionContext) LT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLT, i)
}

func (s *ExpressionContext) AllGT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserGT)
}

func (s *ExpressionContext) GT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGT, i)
}

func (s *ExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLE, 0)
}

func (s *ExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserGE, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEQUAL, 0)
}

func (s *ExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNOTEQUAL, 0)
}

func (s *ExpressionContext) BITAND() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBITAND, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCARET, 0)
}

func (s *ExpressionContext) BITOR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBITOR, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOR, 0)
}

func (s *ExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *ExpressionContext) QUESTION() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserQUESTION, 0)
}

func (s *ExpressionContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserASSIGN, 0)
}

func (s *ExpressionContext) ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserADD_ASSIGN, 0)
}

func (s *ExpressionContext) SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUB_ASSIGN, 0)
}

func (s *ExpressionContext) MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMUL_ASSIGN, 0)
}

func (s *ExpressionContext) DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDIV_ASSIGN, 0)
}

func (s *ExpressionContext) AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserAND_ASSIGN, 0)
}

func (s *ExpressionContext) OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOR_ASSIGN, 0)
}

func (s *ExpressionContext) XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserXOR_ASSIGN, 0)
}

func (s *ExpressionContext) RSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) URSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserURSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) LSHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLSHIFT_ASSIGN, 0)
}

func (s *ExpressionContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserMOD_ASSIGN, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *ExpressionContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *ExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, 0)
}

func (s *ExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, 0)
}

func (s *ExpressionContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ExpressionContext) INSTANCEOF() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserINSTANCEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PlayScriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, PlayScriptParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(127)
			p.Primary()
		}

	case 2:
		{
			p.SetState(128)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(129)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-85)&-(0x1f+1)) == 0 && ((1<<uint((_la-85)))&((1<<(PlayScriptParserINC-85))|(1<<(PlayScriptParserDEC-85))|(1<<(PlayScriptParserADD-85))|(1<<(PlayScriptParserSUB-85)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(130)
			p.expression(15)
		}

	case 4:
		{
			p.SetState(131)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).prefix = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PlayScriptParserBANG || _la == PlayScriptParserTILDE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).prefix = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(132)
			p.expression(14)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(136)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-89)&-(0x1f+1)) == 0 && ((1<<uint((_la-89)))&((1<<(PlayScriptParserMUL-89))|(1<<(PlayScriptParserDIV-89))|(1<<(PlayScriptParserMOD-89)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(137)
					p.expression(14)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(139)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserADD || _la == PlayScriptParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(140)
					p.expression(13)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(149)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(142)
						p.Match(PlayScriptParserLT)
					}
					{
						p.SetState(143)
						p.Match(PlayScriptParserLT)
					}

				case 2:
					{
						p.SetState(144)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(145)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(146)
						p.Match(PlayScriptParserGT)
					}

				case 3:
					{
						p.SetState(147)
						p.Match(PlayScriptParserGT)
					}
					{
						p.SetState(148)
						p.Match(PlayScriptParserGT)
					}

				}
				{
					p.SetState(151)
					p.expression(12)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(153)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-73)&-(0x1f+1)) == 0 && ((1<<uint((_la-73)))&((1<<(PlayScriptParserGT-73))|(1<<(PlayScriptParserLT-73))|(1<<(PlayScriptParserLE-73))|(1<<(PlayScriptParserGE-73)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(154)
					p.expression(11)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(156)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserEQUAL || _la == PlayScriptParserNOTEQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(157)
					p.expression(9)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(159)

					var _m = p.Match(PlayScriptParserBITAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(160)
					p.expression(8)
				}

			case 7:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(162)

					var _m = p.Match(PlayScriptParserCARET)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(163)
					p.expression(7)
				}

			case 8:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(164)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(165)

					var _m = p.Match(PlayScriptParserBITOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(166)
					p.expression(6)
				}

			case 9:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(167)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(168)

					var _m = p.Match(PlayScriptParserAND)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(169)
					p.expression(5)
				}

			case 10:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(170)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(171)

					var _m = p.Match(PlayScriptParserOR)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(172)
					p.expression(4)
				}

			case 11:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(173)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(174)

					var _m = p.Match(PlayScriptParserQUESTION)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(175)
					p.expression(0)
				}
				{
					p.SetState(176)
					p.Match(PlayScriptParserCOLON)
				}
				{
					p.SetState(177)
					p.expression(3)
				}

			case 12:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(180)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).bop = _lt

					_la = p.GetTokenStream().LA(1)

					if !((((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(PlayScriptParserASSIGN-72))|(1<<(PlayScriptParserADD_ASSIGN-72))|(1<<(PlayScriptParserSUB_ASSIGN-72))|(1<<(PlayScriptParserMUL_ASSIGN-72))|(1<<(PlayScriptParserDIV_ASSIGN-72))|(1<<(PlayScriptParserAND_ASSIGN-72))|(1<<(PlayScriptParserOR_ASSIGN-72))|(1<<(PlayScriptParserXOR_ASSIGN-72))|(1<<(PlayScriptParserMOD_ASSIGN-72))|(1<<(PlayScriptParserLSHIFT_ASSIGN-72)))) != 0) || _la == PlayScriptParserRSHIFT_ASSIGN || _la == PlayScriptParserURSHIFT_ASSIGN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).bop = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(181)
					p.expression(1)
				}

			case 13:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(182)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(183)

					var _m = p.Match(PlayScriptParserDOT)

					localctx.(*ExpressionContext).bop = _m
				}
				p.SetState(187)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(184)
						p.Match(PlayScriptParserIDENTIFIER)
					}

				case 2:
					{
						p.SetState(185)
						p.FunctionCall()
					}

				case 3:
					{
						p.SetState(186)
						p.Match(PlayScriptParserTHIS)
					}

				}

			case 14:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(190)
					p.Match(PlayScriptParserLBRACK)
				}
				{
					p.SetState(191)
					p.expression(0)
				}
				{
					p.SetState(192)
					p.Match(PlayScriptParserRBRACK)
				}

			case 15:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(195)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).postfix = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PlayScriptParserINC || _la == PlayScriptParserDEC) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).postfix = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 16:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PlayScriptParserRULE_expression)
				p.SetState(196)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(197)

					var _m = p.Match(PlayScriptParserINSTANCEOF)

					localctx.(*ExpressionContext).bop = _m
				}
				{
					p.SetState(198)
					p.TypeType()
				}

			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PlayScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.expression(0)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(205)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(206)
			p.expression(0)
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FunctionCallContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FunctionCallContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *FunctionCallContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUPER, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PlayScriptParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(PlayScriptParserIDENTIFIER)
		}
		{
			p.SetState(213)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(214)
				p.ExpressionList()
			}

		}
		{
			p.SetState(217)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(PlayScriptParserTHIS)
		}
		{
			p.SetState(219)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(220)
				p.ExpressionList()
			}

		}
		{
			p.SetState(223)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.Match(PlayScriptParserSUPER)
		}
		{
			p.SetState(225)
			p.Match(PlayScriptParserLPAREN)
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(226)
				p.ExpressionList()
			}

		}
		{
			p.SetState(229)
			p.Match(PlayScriptParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *PrimaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *PrimaryContext) THIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHIS, 0)
}

func (s *PrimaryContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSUPER, 0)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PlayScriptParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(PlayScriptParserLPAREN)
		}
		{
			p.SetState(233)
			p.expression(0)
		}
		{
			p.SetState(234)
			p.Match(PlayScriptParserRPAREN)
		}

	case PlayScriptParserTHIS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(PlayScriptParserTHIS)
		}

	case PlayScriptParserSUPER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.Match(PlayScriptParserSUPER)
		}

	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL, PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL, PlayScriptParserBOOL_LITERAL, PlayScriptParserCHAR_LITERAL, PlayScriptParserSTRING_LITERAL, PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)
			p.Literal()
		}

	case PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(239)
			p.Match(PlayScriptParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntLiteralContext is an interface to support dynamic dispatch.
type IIntLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntLiteralContext differentiates from other interfaces.
	IsIntLiteralContext()
}

type IntLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLiteralContext() *IntLiteralContext {
	var p = new(IntLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_intLiteral
	return p
}

func (*IntLiteralContext) IsIntLiteralContext() {}

func NewIntLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLiteralContext {
	var p = new(IntLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_intLiteral

	return p
}

func (s *IntLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDECIMAL_LITERAL, 0)
}

func (s *IntLiteralContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_LITERAL, 0)
}

func (s *IntLiteralContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserOCT_LITERAL, 0)
}

func (s *IntLiteralContext) BINARY_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBINARY_LITERAL, 0)
}

func (s *IntLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterIntLiteral(s)
	}
}

func (s *IntLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitIntLiteral(s)
	}
}

func (s *IntLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitIntLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) IntLiteral() (localctx IIntLiteralContext) {
	localctx = NewIntLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PlayScriptParserRULE_intLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(PlayScriptParserDECIMAL_LITERAL-53))|(1<<(PlayScriptParserHEX_LITERAL-53))|(1<<(PlayScriptParserOCT_LITERAL-53))|(1<<(PlayScriptParserBINARY_LITERAL-53)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) HEX_FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserHEX_FLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFloatLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FloatLiteral() (localctx IFloatLiteralContext) {
	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PlayScriptParserRULE_floatLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PlayScriptParserFLOAT_LITERAL || _la == PlayScriptParserHEX_FLOAT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntLiteral() IIntLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntLiteralContext)
}

func (s *LiteralContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *LiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCHAR_LITERAL, 0)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserNULL_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PlayScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(252)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.IntLiteral()
		}

	case PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.FloatLiteral()
		}

	case PlayScriptParserCHAR_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(248)
			p.Match(PlayScriptParserCHAR_LITERAL)
		}

	case PlayScriptParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(249)
			p.Match(PlayScriptParserSTRING_LITERAL)
		}

	case PlayScriptParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(250)
			p.Match(PlayScriptParserBOOL_LITERAL)
		}

	case PlayScriptParserNULL_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(251)
			p.Match(PlayScriptParserNULL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *BlockContext) BlockStatements() IBlockStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementsContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PlayScriptParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(PlayScriptParserLBRACE)
	}
	{
		p.SetState(255)
		p.BlockStatements()
	}
	{
		p.SetState(256)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *BlockStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *BlockStatementContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlockStatement(s)
	}
}

func (s *BlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlockStatement(s)
	}
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PlayScriptParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.VariableDeclarators()
		}
		{
			p.SetState(259)
			p.Match(PlayScriptParserSEMI)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(262)
			p.FunctionDeclaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(263)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IBlockStatementsContext is an interface to support dynamic dispatch.
type IBlockStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementsContext differentiates from other interfaces.
	IsBlockStatementsContext()
}

type BlockStatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementsContext() *BlockStatementsContext {
	var p = new(BlockStatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_blockStatements
	return p
}

func (*BlockStatementsContext) IsBlockStatementsContext() {}

func NewBlockStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementsContext {
	var p = new(BlockStatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_blockStatements

	return p
}

func (s *BlockStatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementsContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockStatementsContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockStatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterBlockStatements(s)
	}
}

func (s *BlockStatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitBlockStatements(s)
	}
}

func (s *BlockStatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitBlockStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) BlockStatements() (localctx IBlockStatementsContext) {
	localctx = NewBlockStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PlayScriptParserRULE_blockStatements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBREAK)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserCONTINUE)|(1<<PlayScriptParserDO)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserFOR)|(1<<PlayScriptParserIF)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(PlayScriptParserRETURN-36))|(1<<(PlayScriptParserSHORT-36))|(1<<(PlayScriptParserSUPER-36))|(1<<(PlayScriptParserSWITCH-36))|(1<<(PlayScriptParserTHIS-36))|(1<<(PlayScriptParserVOID-36))|(1<<(PlayScriptParserWHILE-36))|(1<<(PlayScriptParserFUNCTION-36))|(1<<(PlayScriptParserSTRING-36))|(1<<(PlayScriptParserDECIMAL_LITERAL-36))|(1<<(PlayScriptParserHEX_LITERAL-36))|(1<<(PlayScriptParserOCT_LITERAL-36))|(1<<(PlayScriptParserBINARY_LITERAL-36))|(1<<(PlayScriptParserFLOAT_LITERAL-36))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-36))|(1<<(PlayScriptParserBOOL_LITERAL-36))|(1<<(PlayScriptParserCHAR_LITERAL-36))|(1<<(PlayScriptParserSTRING_LITERAL-36))|(1<<(PlayScriptParserNULL_LITERAL-36))|(1<<(PlayScriptParserLPAREN-36))|(1<<(PlayScriptParserLBRACE-36)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(PlayScriptParserSEMI-69))|(1<<(PlayScriptParserBANG-69))|(1<<(PlayScriptParserTILDE-69))|(1<<(PlayScriptParserINC-69))|(1<<(PlayScriptParserDEC-69))|(1<<(PlayScriptParserADD-69))|(1<<(PlayScriptParserSUB-69)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(266)
			p.BlockStatement()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionBodyContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PlayScriptParserRULE_functionBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(274)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Block()
		}

	case PlayScriptParserSEMI:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Match(PlayScriptParserSEMI)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserIDENTIFIER)
}

func (s *QualifiedNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, i)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDOT, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PlayScriptParserRULE_qualifiedName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserDOT {
		{
			p.SetState(277)
			p.Match(PlayScriptParserDOT)
		}
		{
			p.SetState(278)
			p.Match(PlayScriptParserIDENTIFIER)
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQualifiedNameListContext is an interface to support dynamic dispatch.
type IQualifiedNameListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameListContext differentiates from other interfaces.
	IsQualifiedNameListContext()
}

type QualifiedNameListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameListContext() *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_qualifiedNameList
	return p
}

func (*QualifiedNameListContext) IsQualifiedNameListContext() {}

func NewQualifiedNameListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameListContext {
	var p = new(QualifiedNameListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_qualifiedNameList

	return p
}

func (s *QualifiedNameListContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameListContext) AllQualifiedName() []IQualifiedNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem())
	var tst = make([]IQualifiedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualifiedNameContext)
		}
	}

	return tst
}

func (s *QualifiedNameListContext) QualifiedName(i int) IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *QualifiedNameListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *QualifiedNameListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *QualifiedNameListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitQualifiedNameList(s)
	}
}

func (s *QualifiedNameListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitQualifiedNameList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) QualifiedNameList() (localctx IQualifiedNameListContext) {
	localctx = NewQualifiedNameListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PlayScriptParserRULE_qualifiedNameList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.QualifiedName()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(285)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(286)
			p.QualifiedName()
		}

		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) TypeTypeOrVoid() ITypeTypeOrVoidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeOrVoidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeOrVoidContext)
}

func (s *FunctionDeclarationContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *FunctionDeclarationContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *FunctionDeclarationContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *FunctionDeclarationContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *FunctionDeclarationContext) THROWS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserTHROWS, 0)
}

func (s *FunctionDeclarationContext) QualifiedNameList() IQualifiedNameListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PlayScriptParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(292)
			p.TypeTypeOrVoid()
		}

	}
	{
		p.SetState(295)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	{
		p.SetState(296)
		p.FormalParameters()
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserLBRACK {
		{
			p.SetState(297)
			p.Match(PlayScriptParserLBRACK)
		}
		{
			p.SetState(298)
			p.Match(PlayScriptParserRBRACK)
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserTHROWS {
		{
			p.SetState(304)
			p.Match(PlayScriptParserTHROWS)
		}
		{
			p.SetState(305)
			p.QualifiedNameList()
		}

	}
	{
		p.SetState(308)
		p.FunctionBody()
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *FormalParametersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *FormalParametersContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameters(s)
	}
}

func (s *FormalParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameters(s)
	}
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PlayScriptParserRULE_formalParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(PlayScriptParserLPAREN)
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFINAL)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(311)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(314)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllFormalParameter() []IFormalParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem())
	var tst = make([]IFormalParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalParameterContext)
		}
	}

	return tst
}

func (s *FormalParameterListContext) FormalParameter(i int) IFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterContext)
}

func (s *FormalParameterListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *FormalParameterListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *FormalParameterListContext) LastFormalParameter() ILastFormalParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILastFormalParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILastFormalParameterContext)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PlayScriptParserRULE_formalParameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.FormalParameter()
		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(317)
					p.Match(PlayScriptParserCOMMA)
				}
				{
					p.SetState(318)
					p.FormalParameter()
				}

			}
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserCOMMA {
			{
				p.SetState(324)
				p.Match(PlayScriptParserCOMMA)
			}
			{
				p.SetState(325)
				p.LastFormalParameter()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)
			p.LastFormalParameter()
		}

	}

	return localctx
}

// IFormalParameterContext is an interface to support dynamic dispatch.
type IFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterContext differentiates from other interfaces.
	IsFormalParameterContext()
}

type FormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterContext() *FormalParameterContext {
	var p = new(FormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_formalParameter
	return p
}

func (*FormalParameterContext) IsFormalParameterContext() {}

func NewFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterContext {
	var p = new(FormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_formalParameter

	return p
}

func (s *FormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *FormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *FormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *FormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *FormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFormalParameter(s)
	}
}

func (s *FormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFormalParameter(s)
	}
}

func (s *FormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FormalParameter() (localctx IFormalParameterContext) {
	localctx = NewFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PlayScriptParserRULE_formalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserFINAL {
		{
			p.SetState(331)
			p.VariableModifier()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(337)
		p.TypeType()
	}
	{
		p.SetState(338)
		p.VariableDeclaratorId()
	}

	return localctx
}

// ILastFormalParameterContext is an interface to support dynamic dispatch.
type ILastFormalParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLastFormalParameterContext differentiates from other interfaces.
	IsLastFormalParameterContext()
}

type LastFormalParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLastFormalParameterContext() *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_lastFormalParameter
	return p
}

func (*LastFormalParameterContext) IsLastFormalParameterContext() {}

func NewLastFormalParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LastFormalParameterContext {
	var p = new(LastFormalParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_lastFormalParameter

	return p
}

func (s *LastFormalParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *LastFormalParameterContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *LastFormalParameterContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserELLIPSIS, 0)
}

func (s *LastFormalParameterContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *LastFormalParameterContext) AllVariableModifier() []IVariableModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem())
	var tst = make([]IVariableModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableModifierContext)
		}
	}

	return tst
}

func (s *LastFormalParameterContext) VariableModifier(i int) IVariableModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableModifierContext)
}

func (s *LastFormalParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LastFormalParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LastFormalParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitLastFormalParameter(s)
	}
}

func (s *LastFormalParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitLastFormalParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) LastFormalParameter() (localctx ILastFormalParameterContext) {
	localctx = NewLastFormalParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PlayScriptParserRULE_lastFormalParameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserFINAL {
		{
			p.SetState(340)
			p.VariableModifier()
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(346)
		p.TypeType()
	}
	{
		p.SetState(347)
		p.Match(PlayScriptParserELLIPSIS)
	}
	{
		p.SetState(348)
		p.VariableDeclaratorId()
	}

	return localctx
}

// IVariableModifierContext is an interface to support dynamic dispatch.
type IVariableModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableModifierContext differentiates from other interfaces.
	IsVariableModifierContext()
}

type VariableModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableModifierContext() *VariableModifierContext {
	var p = new(VariableModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableModifier
	return p
}

func (*VariableModifierContext) IsVariableModifierContext() {}

func NewVariableModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableModifierContext {
	var p = new(VariableModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableModifier

	return p
}

func (s *VariableModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableModifierContext) FINAL() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFINAL, 0)
}

func (s *VariableModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableModifier(s)
	}
}

func (s *VariableModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableModifier(s)
	}
}

func (s *VariableModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableModifier() (localctx IVariableModifierContext) {
	localctx = NewVariableModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PlayScriptParserRULE_variableModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(PlayScriptParserFINAL)
	}

	return localctx
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classDeclaration
	return p
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) CLASS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCLASS, 0)
}

func (s *ClassDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *ClassDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ClassDeclarationContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserEXTENDS, 0)
}

func (s *ClassDeclarationContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *ClassDeclarationContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIMPLEMENTS, 0)
}

func (s *ClassDeclarationContext) TypeList() ITypeListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeListContext)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PlayScriptParserRULE_classDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(PlayScriptParserCLASS)
	}
	{
		p.SetState(353)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserEXTENDS {
		{
			p.SetState(354)
			p.Match(PlayScriptParserEXTENDS)
		}
		{
			p.SetState(355)
			p.TypeType()
		}

	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PlayScriptParserIMPLEMENTS {
		{
			p.SetState(358)
			p.Match(PlayScriptParserIMPLEMENTS)
		}
		{
			p.SetState(359)
			p.TypeList()
		}

	}
	{
		p.SetState(362)
		p.ClassBody()
	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *ClassBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *ClassBodyContext) AllClassBodyDeclaration() []IClassBodyDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem())
	var tst = make([]IClassBodyDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBodyDeclarationContext)
		}
	}

	return tst
}

func (s *ClassBodyContext) ClassBodyDeclaration(i int) IClassBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBodyDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (s *ClassBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassBody() (localctx IClassBodyContext) {
	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PlayScriptParserRULE_classBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.Match(PlayScriptParserLBRACE)
	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserVOID-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37)))) != 0) || _la == PlayScriptParserSEMI || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(365)
			p.ClassBodyDeclaration()
		}

		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(371)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IClassBodyDeclarationContext is an interface to support dynamic dispatch.
type IClassBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyDeclarationContext differentiates from other interfaces.
	IsClassBodyDeclarationContext()
}

type ClassBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyDeclarationContext() *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_classBodyDeclaration
	return p
}

func (*ClassBodyDeclarationContext) IsClassBodyDeclarationContext() {}

func NewClassBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_classBodyDeclaration

	return p
}

func (s *ClassBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *ClassBodyDeclarationContext) MemberDecalaration() IMemberDecalarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDecalarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberDecalarationContext)
}

func (s *ClassBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitClassBodyDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ClassBodyDeclaration() (localctx IClassBodyDeclarationContext) {
	localctx = NewClassBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PlayScriptParserRULE_classBodyDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserSEMI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(373)
			p.Match(PlayScriptParserSEMI)
		}

	case PlayScriptParserBOOLEAN, PlayScriptParserBYTE, PlayScriptParserCHAR, PlayScriptParserCLASS, PlayScriptParserDOUBLE, PlayScriptParserFLOAT, PlayScriptParserINT, PlayScriptParserLONG, PlayScriptParserSHORT, PlayScriptParserVOID, PlayScriptParserFUNCTION, PlayScriptParserSTRING, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(374)
			p.MemberDecalaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMemberDecalarationContext is an interface to support dynamic dispatch.
type IMemberDecalarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDecalarationContext differentiates from other interfaces.
	IsMemberDecalarationContext()
}

type MemberDecalarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDecalarationContext() *MemberDecalarationContext {
	var p = new(MemberDecalarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_memberDecalaration
	return p
}

func (*MemberDecalarationContext) IsMemberDecalarationContext() {}

func NewMemberDecalarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDecalarationContext {
	var p = new(MemberDecalarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_memberDecalaration

	return p
}

func (s *MemberDecalarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDecalarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *MemberDecalarationContext) FieldDeclaration() IFieldDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *MemberDecalarationContext) ClassDeclaration() IClassDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *MemberDecalarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDecalarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberDecalarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterMemberDecalaration(s)
	}
}

func (s *MemberDecalarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitMemberDecalaration(s)
	}
}

func (s *MemberDecalarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitMemberDecalaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) MemberDecalaration() (localctx IMemberDecalarationContext) {
	localctx = NewMemberDecalarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PlayScriptParserRULE_memberDecalaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.FunctionDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)
			p.FieldDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(379)
			p.ClassDeclaration()
		}

	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *FieldDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PlayScriptParserRULE_fieldDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.VariableDeclarators()
	}
	{
		p.SetState(383)
		p.Match(PlayScriptParserSEMI)
	}

	return localctx
}

// IParExpressionContext is an interface to support dynamic dispatch.
type IParExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParExpressionContext differentiates from other interfaces.
	IsParExpressionContext()
}

type ParExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParExpressionContext() *ParExpressionContext {
	var p = new(ParExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_parExpression
	return p
}

func (*ParExpressionContext) IsParExpressionContext() {}

func NewParExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParExpressionContext {
	var p = new(ParExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_parExpression

	return p
}

func (s *ParExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *ParExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *ParExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterParExpression(s)
	}
}

func (s *ParExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitParExpression(s)
	}
}

func (s *ParExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitParExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ParExpression() (localctx IParExpressionContext) {
	localctx = NewParExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PlayScriptParserRULE_parExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(385)
		p.Match(PlayScriptParserLPAREN)
	}
	{
		p.SetState(386)
		p.expression(0)
	}
	{
		p.SetState(387)
		p.Match(PlayScriptParserRPAREN)
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetForUpdate returns the forUpdate rule contexts.
	GetForUpdate() IExpressionListContext

	// SetForUpdate sets the forUpdate rule contexts.
	SetForUpdate(IExpressionListContext)

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	forUpdate IExpressionListContext
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) GetForUpdate() IExpressionListContext { return s.forUpdate }

func (s *ForControlContext) SetForUpdate(v IExpressionListContext) { s.forUpdate = v }

func (s *ForControlContext) EnhanceForControl() IEnhanceForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnhanceForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnhanceForControlContext)
}

func (s *ForControlContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserSEMI)
}

func (s *ForControlContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, i)
}

func (s *ForControlContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterForControl(s)
	}
}

func (s *ForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitForControl(s)
	}
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PlayScriptParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.EnhanceForControl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(PlayScriptParserSHORT-37))|(1<<(PlayScriptParserSUPER-37))|(1<<(PlayScriptParserTHIS-37))|(1<<(PlayScriptParserFUNCTION-37))|(1<<(PlayScriptParserSTRING-37))|(1<<(PlayScriptParserDECIMAL_LITERAL-37))|(1<<(PlayScriptParserHEX_LITERAL-37))|(1<<(PlayScriptParserOCT_LITERAL-37))|(1<<(PlayScriptParserBINARY_LITERAL-37))|(1<<(PlayScriptParserFLOAT_LITERAL-37))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-37))|(1<<(PlayScriptParserBOOL_LITERAL-37))|(1<<(PlayScriptParserCHAR_LITERAL-37))|(1<<(PlayScriptParserSTRING_LITERAL-37))|(1<<(PlayScriptParserNULL_LITERAL-37))|(1<<(PlayScriptParserLPAREN-37)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(390)
				p.ForInit()
			}

		}
		{
			p.SetState(393)
			p.Match(PlayScriptParserSEMI)
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(394)
				p.expression(0)
			}

		}
		{
			p.SetState(397)
			p.Match(PlayScriptParserSEMI)
		}
		{
			p.SetState(398)

			var _x = p.ExpressionList()

			localctx.(*ForControlContext).forUpdate = _x
		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterForInit(s)
	}
}

func (s *ForInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitForInit(s)
	}
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PlayScriptParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.VariableDeclarators()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.ExpressionList()
		}

	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *VariableDeclaratorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserASSIGN, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclarator(s)
	}
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PlayScriptParserRULE_variableDeclarator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(405)
		p.VariableDeclaratorId()
	}

	{
		p.SetState(406)
		p.Match(PlayScriptParserASSIGN)
	}
	{
		p.SetState(407)
		p.VariableInitializer()
	}

	return localctx
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_arrayInitializer
	return p
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *ArrayInitializerContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem())
	var tst = make([]IVariableInitializerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableInitializerContext)
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *ArrayInitializerContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *ArrayInitializerContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitArrayInitializer(s)
	}
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PlayScriptParserRULE_arrayInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.Match(PlayScriptParserLBRACE)
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40))|(1<<(PlayScriptParserLBRACE-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(410)
			p.VariableInitializer()
		}
		p.SetState(415)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(411)
					p.Match(PlayScriptParserCOMMA)
				}
				{
					p.SetState(412)
					p.VariableInitializer()
				}

			}
			p.SetState(417)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserCOMMA {
			{
				p.SetState(418)
				p.Match(PlayScriptParserCOMMA)
			}

		}

	}
	{
		p.SetState(423)
		p.Match(PlayScriptParserRBRACE)
	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableInitializer(s)
	}
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PlayScriptParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(427)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.ArrayInitializer()
		}

	case PlayScriptParserSUPER, PlayScriptParserTHIS, PlayScriptParserDECIMAL_LITERAL, PlayScriptParserHEX_LITERAL, PlayScriptParserOCT_LITERAL, PlayScriptParserBINARY_LITERAL, PlayScriptParserFLOAT_LITERAL, PlayScriptParserHEX_FLOAT_LITERAL, PlayScriptParserBOOL_LITERAL, PlayScriptParserCHAR_LITERAL, PlayScriptParserSTRING_LITERAL, PlayScriptParserNULL_LITERAL, PlayScriptParserLPAREN, PlayScriptParserBANG, PlayScriptParserTILDE, PlayScriptParserINC, PlayScriptParserDEC, PlayScriptParserADD, PlayScriptParserSUB, PlayScriptParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(426)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserCOMMA)
}

func (s *VariableDeclaratorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOMMA, i)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclarators(s)
	}
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PlayScriptParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(429)
		p.TypeType()
	}
	{
		p.SetState(430)
		p.VariableDeclarator()
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserCOMMA {
		{
			p.SetState(431)
			p.Match(PlayScriptParserCOMMA)
		}
		{
			p.SetState(432)
			p.VariableDeclarator()
		}

		p.SetState(437)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IEnhanceForControlContext is an interface to support dynamic dispatch.
type IEnhanceForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnhanceForControlContext differentiates from other interfaces.
	IsEnhanceForControlContext()
}

type EnhanceForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnhanceForControlContext() *EnhanceForControlContext {
	var p = new(EnhanceForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_enhanceForControl
	return p
}

func (*EnhanceForControlContext) IsEnhanceForControlContext() {}

func NewEnhanceForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnhanceForControlContext {
	var p = new(EnhanceForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_enhanceForControl

	return p
}

func (s *EnhanceForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *EnhanceForControlContext) TypeType() ITypeTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeTypeContext)
}

func (s *EnhanceForControlContext) VariableDeclaratorId() IVariableDeclaratorIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorIdContext)
}

func (s *EnhanceForControlContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *EnhanceForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EnhanceForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnhanceForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnhanceForControlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterEnhanceForControl(s)
	}
}

func (s *EnhanceForControlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitEnhanceForControl(s)
	}
}

func (s *EnhanceForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitEnhanceForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) EnhanceForControl() (localctx IEnhanceForControlContext) {
	localctx = NewEnhanceForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PlayScriptParserRULE_enhanceForControl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		p.TypeType()
	}
	{
		p.SetState(439)
		p.VariableDeclaratorId()
	}
	{
		p.SetState(440)
		p.Match(PlayScriptParserCOLON)
	}
	{
		p.SetState(441)
		p.expression(0)
	}

	return localctx
}

// IVariableDeclaratorIdContext is an interface to support dynamic dispatch.
type IVariableDeclaratorIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorIdContext differentiates from other interfaces.
	IsVariableDeclaratorIdContext()
}

type VariableDeclaratorIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorIdContext() *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_variableDeclaratorId
	return p
}

func (*VariableDeclaratorIdContext) IsVariableDeclaratorIdContext() {}

func NewVariableDeclaratorIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorIdContext {
	var p = new(VariableDeclaratorIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_variableDeclaratorId

	return p
}

func (s *VariableDeclaratorIdContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorIdContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *VariableDeclaratorIdContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserLBRACK)
}

func (s *VariableDeclaratorIdContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACK, i)
}

func (s *VariableDeclaratorIdContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(PlayScriptParserRBRACK)
}

func (s *VariableDeclaratorIdContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACK, i)
}

func (s *VariableDeclaratorIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitVariableDeclaratorId(s)
	}
}

func (s *VariableDeclaratorIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitVariableDeclaratorId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) VariableDeclaratorId() (localctx IVariableDeclaratorIdContext) {
	localctx = NewVariableDeclaratorIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PlayScriptParserRULE_variableDeclaratorId)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
		p.Match(PlayScriptParserIDENTIFIER)
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PlayScriptParserLBRACK {
		{
			p.SetState(444)
			p.Match(PlayScriptParserLBRACK)
		}
		{
			p.SetState(445)
			p.Match(PlayScriptParserRBRACK)
		}

		p.SetState(450)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchBlockStatementGroupContext is an interface to support dynamic dispatch.
type ISwitchBlockStatementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockStatementGroupContext differentiates from other interfaces.
	IsSwitchBlockStatementGroupContext()
}

type SwitchBlockStatementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockStatementGroupContext() *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_switchBlockStatementGroup
	return p
}

func (*SwitchBlockStatementGroupContext) IsSwitchBlockStatementGroupContext() {}

func NewSwitchBlockStatementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_switchBlockStatementGroup

	return p
}

func (s *SwitchBlockStatementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockStatementGroupContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockStatementGroupContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *SwitchBlockStatementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockStatementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockStatementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitSwitchBlockStatementGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) SwitchBlockStatementGroup() (localctx ISwitchBlockStatementGroupContext) {
	localctx = NewSwitchBlockStatementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PlayScriptParserRULE_switchBlockStatementGroup)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PlayScriptParserCASE || _la == PlayScriptParserDEFAULT {
		{
			p.SetState(451)
			p.SwitchLabel()
		}

		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PlayScriptParserBOOLEAN)|(1<<PlayScriptParserBREAK)|(1<<PlayScriptParserBYTE)|(1<<PlayScriptParserCHAR)|(1<<PlayScriptParserCLASS)|(1<<PlayScriptParserCONTINUE)|(1<<PlayScriptParserDO)|(1<<PlayScriptParserDOUBLE)|(1<<PlayScriptParserFLOAT)|(1<<PlayScriptParserFOR)|(1<<PlayScriptParserIF)|(1<<PlayScriptParserINT)|(1<<PlayScriptParserLONG))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(PlayScriptParserRETURN-36))|(1<<(PlayScriptParserSHORT-36))|(1<<(PlayScriptParserSUPER-36))|(1<<(PlayScriptParserSWITCH-36))|(1<<(PlayScriptParserTHIS-36))|(1<<(PlayScriptParserVOID-36))|(1<<(PlayScriptParserWHILE-36))|(1<<(PlayScriptParserFUNCTION-36))|(1<<(PlayScriptParserSTRING-36))|(1<<(PlayScriptParserDECIMAL_LITERAL-36))|(1<<(PlayScriptParserHEX_LITERAL-36))|(1<<(PlayScriptParserOCT_LITERAL-36))|(1<<(PlayScriptParserBINARY_LITERAL-36))|(1<<(PlayScriptParserFLOAT_LITERAL-36))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-36))|(1<<(PlayScriptParserBOOL_LITERAL-36))|(1<<(PlayScriptParserCHAR_LITERAL-36))|(1<<(PlayScriptParserSTRING_LITERAL-36))|(1<<(PlayScriptParserNULL_LITERAL-36))|(1<<(PlayScriptParserLPAREN-36))|(1<<(PlayScriptParserLBRACE-36)))) != 0) || (((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(PlayScriptParserSEMI-69))|(1<<(PlayScriptParserBANG-69))|(1<<(PlayScriptParserTILDE-69))|(1<<(PlayScriptParserINC-69))|(1<<(PlayScriptParserDEC-69))|(1<<(PlayScriptParserADD-69))|(1<<(PlayScriptParserSUB-69)))) != 0) || _la == PlayScriptParserIDENTIFIER {
		{
			p.SetState(456)
			p.BlockStatement()
		}

		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchLabelContext is an interface to support dynamic dispatch.
type ISwitchLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEnumConstantName returns the enumConstantName token.
	GetEnumConstantName() antlr.Token

	// SetEnumConstantName sets the enumConstantName token.
	SetEnumConstantName(antlr.Token)

	// GetConstantExpression returns the constantExpression rule contexts.
	GetConstantExpression() IExpressionContext

	// SetConstantExpression sets the constantExpression rule contexts.
	SetConstantExpression(IExpressionContext)

	// IsSwitchLabelContext differentiates from other interfaces.
	IsSwitchLabelContext()
}

type SwitchLabelContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	constantExpression IExpressionContext
	enumConstantName   antlr.Token
}

func NewEmptySwitchLabelContext() *SwitchLabelContext {
	var p = new(SwitchLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_switchLabel
	return p
}

func (*SwitchLabelContext) IsSwitchLabelContext() {}

func NewSwitchLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchLabelContext {
	var p = new(SwitchLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_switchLabel

	return p
}

func (s *SwitchLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchLabelContext) GetEnumConstantName() antlr.Token { return s.enumConstantName }

func (s *SwitchLabelContext) SetEnumConstantName(v antlr.Token) { s.enumConstantName = v }

func (s *SwitchLabelContext) GetConstantExpression() IExpressionContext { return s.constantExpression }

func (s *SwitchLabelContext) SetConstantExpression(v IExpressionContext) { s.constantExpression = v }

func (s *SwitchLabelContext) CASE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCASE, 0)
}

func (s *SwitchLabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *SwitchLabelContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchLabelContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *SwitchLabelContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDEFAULT, 0)
}

func (s *SwitchLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitSwitchLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) SwitchLabel() (localctx ISwitchLabelContext) {
	localctx = NewSwitchLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PlayScriptParserRULE_switchLabel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(469)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PlayScriptParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(PlayScriptParserCASE)
		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(462)

				var _x = p.expression(0)

				localctx.(*SwitchLabelContext).constantExpression = _x
			}

		case 2:
			{
				p.SetState(463)

				var _m = p.Match(PlayScriptParserIDENTIFIER)

				localctx.(*SwitchLabelContext).enumConstantName = _m
			}

		}
		{
			p.SetState(466)
			p.Match(PlayScriptParserCOLON)
		}

	case PlayScriptParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(467)
			p.Match(PlayScriptParserDEFAULT)
		}
		{
			p.SetState(468)
			p.Match(PlayScriptParserCOLON)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIdentifierLabel returns the identifierLabel token.
	GetIdentifierLabel() antlr.Token

	// SetIdentifierLabel sets the identifierLabel token.
	SetIdentifierLabel(antlr.Token)

	// GetBlockLabel returns the blockLabel rule contexts.
	GetBlockLabel() IBlockContext

	// GetStatementExpression returns the statementExpression rule contexts.
	GetStatementExpression() IExpressionContext

	// SetBlockLabel sets the blockLabel rule contexts.
	SetBlockLabel(IBlockContext)

	// SetStatementExpression sets the statementExpression rule contexts.
	SetStatementExpression(IExpressionContext)

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	blockLabel          IBlockContext
	statementExpression IExpressionContext
	identifierLabel     antlr.Token
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlayScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlayScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) GetIdentifierLabel() antlr.Token { return s.identifierLabel }

func (s *StatementContext) SetIdentifierLabel(v antlr.Token) { s.identifierLabel = v }

func (s *StatementContext) GetBlockLabel() IBlockContext { return s.blockLabel }

func (s *StatementContext) GetStatementExpression() IExpressionContext { return s.statementExpression }

func (s *StatementContext) SetBlockLabel(v IBlockContext) { s.blockLabel = v }

func (s *StatementContext) SetStatementExpression(v IExpressionContext) { s.statementExpression = v }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IF() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIF, 0)
}

func (s *StatementContext) ParExpression() IParExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParExpressionContext)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserELSE, 0)
}

func (s *StatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserFOR, 0)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLPAREN, 0)
}

func (s *StatementContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *StatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRPAREN, 0)
}

func (s *StatementContext) WHILE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserWHILE, 0)
}

func (s *StatementContext) DO() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserDO, 0)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSEMI, 0)
}

func (s *StatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserSWITCH, 0)
}

func (s *StatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserLBRACE, 0)
}

func (s *StatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRBRACE, 0)
}

func (s *StatementContext) AllSwitchBlockStatementGroup() []ISwitchBlockStatementGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem())
	var tst = make([]ISwitchBlockStatementGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchBlockStatementGroupContext)
		}
	}

	return tst
}

func (s *StatementContext) SwitchBlockStatementGroup(i int) ISwitchBlockStatementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockStatementGroupContext)
}

func (s *StatementContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *StatementContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserRETURN, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserBREAK, 0)
}

func (s *StatementContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserIDENTIFIER, 0)
}

func (s *StatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCONTINUE, 0)
}

func (s *StatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PlayScriptParserCOLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlayScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PlayScriptVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PlayScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PlayScriptParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(471)

			var _x = p.Block()

			localctx.(*StatementContext).blockLabel = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.Match(PlayScriptParserIF)
		}
		{
			p.SetState(473)
			p.ParExpression()
		}
		{
			p.SetState(474)
			p.Statement()
		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(475)
				p.Match(PlayScriptParserELSE)
			}
			{
				p.SetState(476)
				p.Statement()
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.Match(PlayScriptParserFOR)
		}
		{
			p.SetState(480)
			p.Match(PlayScriptParserLPAREN)
		}
		{
			p.SetState(481)
			p.ForControl()
		}
		{
			p.SetState(482)
			p.Match(PlayScriptParserRPAREN)
		}
		{
			p.SetState(483)
			p.Statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(485)
			p.Match(PlayScriptParserWHILE)
		}
		{
			p.SetState(486)
			p.ParExpression()
		}
		{
			p.SetState(487)
			p.Statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(489)
			p.Match(PlayScriptParserDO)
		}
		{
			p.SetState(490)
			p.Statement()
		}
		{
			p.SetState(491)
			p.Match(PlayScriptParserWHILE)
		}
		{
			p.SetState(492)
			p.ParExpression()
		}
		{
			p.SetState(493)
			p.Match(PlayScriptParserSEMI)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(495)
			p.Match(PlayScriptParserSWITCH)
		}
		{
			p.SetState(496)
			p.ParExpression()
		}
		{
			p.SetState(497)
			p.Match(PlayScriptParserLBRACE)
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(498)
					p.SwitchBlockStatementGroup()
				}

			}
			p.SetState(503)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
		}
		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PlayScriptParserCASE || _la == PlayScriptParserDEFAULT {
			{
				p.SetState(504)
				p.SwitchLabel()
			}

			p.SetState(509)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(510)
			p.Match(PlayScriptParserRBRACE)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(512)
			p.Match(PlayScriptParserRETURN)
		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(PlayScriptParserSUPER-40))|(1<<(PlayScriptParserTHIS-40))|(1<<(PlayScriptParserDECIMAL_LITERAL-40))|(1<<(PlayScriptParserHEX_LITERAL-40))|(1<<(PlayScriptParserOCT_LITERAL-40))|(1<<(PlayScriptParserBINARY_LITERAL-40))|(1<<(PlayScriptParserFLOAT_LITERAL-40))|(1<<(PlayScriptParserHEX_FLOAT_LITERAL-40))|(1<<(PlayScriptParserBOOL_LITERAL-40))|(1<<(PlayScriptParserCHAR_LITERAL-40))|(1<<(PlayScriptParserSTRING_LITERAL-40))|(1<<(PlayScriptParserNULL_LITERAL-40))|(1<<(PlayScriptParserLPAREN-40)))) != 0) || (((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(PlayScriptParserBANG-75))|(1<<(PlayScriptParserTILDE-75))|(1<<(PlayScriptParserINC-75))|(1<<(PlayScriptParserDEC-75))|(1<<(PlayScriptParserADD-75))|(1<<(PlayScriptParserSUB-75)))) != 0) || _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(513)
				p.expression(0)
			}

		}
		{
			p.SetState(516)
			p.Match(PlayScriptParserSEMI)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(517)
			p.Match(PlayScriptParserBREAK)
		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(518)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		{
			p.SetState(521)
			p.Match(PlayScriptParserSEMI)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(522)
			p.Match(PlayScriptParserCONTINUE)
		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PlayScriptParserIDENTIFIER {
			{
				p.SetState(523)
				p.Match(PlayScriptParserIDENTIFIER)
			}

		}
		{
			p.SetState(526)
			p.Match(PlayScriptParserSEMI)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(527)
			p.Match(PlayScriptParserSEMI)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(528)

			var _x = p.expression(0)

			localctx.(*StatementContext).statementExpression = _x
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(529)

			var _m = p.Match(PlayScriptParserIDENTIFIER)

			localctx.(*StatementContext).identifierLabel = _m
		}
		{
			p.SetState(530)
			p.Match(PlayScriptParserCOLON)
		}
		{
			p.SetState(531)
			p.Statement()
		}

	}

	return localctx
}

func (p *PlayScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PlayScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
