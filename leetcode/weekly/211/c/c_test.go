// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[1,3,5,10,15]`, `[1,2,3,4,5]`, 
			`34`,
		},
		{
			`[4,5,6,5]`, `[2,1,2,1]`, 
			`16`,
		},
		{
			`[1,2,3,5]`, `[8,9,10,1]`, 
			`6`,
		},
		// TODO 测试参数的下界和上界
		{
			`[959052,324158,48586,465876,806911,29121,396214,183661,494970,899835,73479,325218,803058,942532,817624,4355,700731,989229,193004,442912,12782,679100,511803,405525,697143,117424,620757,315046,522788,15171,338797,206459,576767,949903,656615,184127,171048,359374,612178,495351,998358,382880,242840,605409,933954,514220,363782,478412,845398,812022,510118,184234,442481,746949,414703,107779,627542,10330,294460,25724,860972,964597,896073,384293,187321,38155,746535,248344,576004,974131,491691,245282,398221,274651,420389,855362,40534,552323,918246,393274,434717,315929,7161,355041,210589,878327,111954,993663,448572,902723,846249,730711,533468,333366,501434,427123,249736,718184,860245,500387,970993,135332,246225,115271,137839,306498,219680,944136,778113,169171,788428,459086,439925,709174,593948,242246,149491,734546,24851,539056,467124,652270,873395,825676,694329,694923,959611,374501,425996,387202,386965,895264,640655,293693,297949,55565,605886,425074,696281,142621,309118,142185,877986,727355,742910,839660,484653,55828,423592,130733,527751,767489,453985,49497,585199,651897,109417,849602,152089,327984,543504,245456,942597,483713,313810,692068,477215,549731,318905,73810,459803,915743,757623,728140,684987,15596,831972,488677,331475,16059,130624,542524,975244,649582,974326,290568,299236,997613,475880,548113,85735,867703,499767,674001,665596,725475,325163,126649,187371,514842,363975,126699,236622,155278,313854,448028,859345,477212,509380,183198,211217,978015,681805,14583,363457,578577,114404,886605,670317,134046,80368,235036,636813,965046,121028,136476,949557,410458,216676,276856,421340,534913,389306,97644,84780,609074,72405,750727,850908,143487,364714,874084,418607,206227,86850,724059,941159,267116,997646,145917,306502,764626,499327,811099,252030,480052,183345,250051,287006,880092,566549,29809,839361,970718,311345,933045,283455,361738,295204,938728,412948,444659,22535,140489,191697,437972,389243,847112,139419,194383,933443,712602,920062,138722,188699,772992,609674,114093,935794,363988,205366,164192,196309,288567,1749,138941,914963,515680,356596,794168,609680,47753,983259,670565,917059,140631,886513,875037,625372,862161,97095,241820,760993,93226,954168,608282,953877,926674,477422,928775,605527,651686,355050,523257,923432,992322,284578,540584,973270,522477,126034,630061,761443,980766,922269,121853,600041,165428,393378,654525,132791,929492,149544,875563,486938,231974,991720,850959,340909,472402,189510,735042,2973,424771,851245,804394,798575,710729,623892,5825,182226,868243,556013,764873,692092,300730,543043,850444,695357,784254,43734,282010,309203,552625,402465,401112,386646,359759,265190,917148,598312,811206,60845,373921,511613,308212,30139,737918,109200,855566,488897,218124,730465,643452,654214,616722,86264,597488,441619,724351,283558,27997,276463,758794,899919,205349,549649,802234,793570,334696,102978,956851,257051,57949,462368,952967,866800,353882,764930,680700,857127,188533,999594,112596,872139,571475,758845,179795,877493,942814,461521,943040,906313,309798,711985,839649,177931,559074,224113,682828,680469,722624,709174,340156,220966,374817,63988,867606,296774,942133,927676,408482,987698,542760,60319,143507,369395,757244,953903,889440,806602,157261,952013,156165,219062,265670,35826]`,
			`[71,684,890,415,393,888,25,826,800,555,794,419,878,151,836,713,747,460,368,479,653,886,806,572,290,676,874,843,592,614,490,703,268,998,238,197,395,38,587,143,887,844,301,430,374,38,664,615,93,216,354,907,352,272,322,516,122,171,820,89,891,935,532,49,703,168,58,361,697,164,750,523,696,268,57,326,791,446,869,580,277,483,835,753,790,752,564,792,113,450,398,722,950,580,372,587,59,561,925,34,167,409,605,791,806,309,176,820,475,110,774,276,576,264,418,326,991,455,959,722,848,170,197,189,401,784,547,946,19,93,554,704,67,838,373,734,46,744,904,312,331,851,692,936,925,152,61,334,86,505,62,478,441,958,765,16,150,47,282,724,133,457,912,923,971,182,613,515,558,988,719,229,974,657,907,992,946,840,759,49,45,162,126,789,709,960,992,517,576,605,567,754,319,131,454,4,399,968,43,269,929,195,320,384,767,187,97,363,651,295,520,345,683,126,865,979,141,995,827,395,82,969,92,510,126,345,395,61,793,530,266,643,512,184,782,381,2,729,161,264,134,332,848,536,603,358,839,24,933,773,973,494,799,997,172,665,193,516,714,187,444,470,679,998,182,621,903,964,849,885,673,727,884,212,669,112,579,345,948,933,583,41,774,654,134,138,575,469,750,90,830,395,430,545,275,746,635,911,339,572,931,836,504,193,32,644,113,106,524,527,327,13,34,813,440,849,754,940,770,310,714,854,570,245,465,251,917,436,318,174,598,893,809,874,851,536,786,470,290,947,195,282,506,121,876,110,570,227,232,303,409,187,232,227,299,478,949,307,581,714,405,520,513,662,884,786,933,208,428,99,675,362,983,10,907,310,653,845,345,917,891,498,739,173,167,731,417,115,327,231,38,334,884,230,942,773,580,613,838,709,110,154,702,556,444,272,840,901,298,595,326,859,693,990,346,459,759,107,674,443,127,713,672,31,717,174,580,166,634,614,511,435,922,406,791,307,127,935,63,179,7,344,650,156,190,406,29,914,929,730,492,773,779,668,38,753,297,92,118,270,585,149,479,569,394,302,641]`,
			`21858209`,
		},
		// [959052,324158,48586,465876,806911,29121,396214,183661,494970,899835,73479,325218,803058,942532,817624,4355,700731,989229,193004,442912,12782,679100,511803,405525,697143,117424,620757,315046,522788,15171,338797,206459,576767,949903,656615,184127,171048,359374,612178,495351,998358,382880,242840,605409,933954,514220,363782,478412,845398,812022,510118,184234,442481,746949,414703,107779,627542,10330,294460,25724,860972,964597,896073,384293,187321,38155,746535,248344,576004,974131,491691,245282,398221,274651,420389,855362,40534,552323,918246,393274,434717,315929,7161,355041,210589,878327,111954,993663,448572,902723,846249,730711,533468,333366,501434,427123,249736,718184,860245,500387,970993,135332,246225,115271,137839,306498,219680,944136,778113,169171,788428,459086,439925,709174,593948,242246,149491,734546,24851,539056,467124,652270,873395,825676,694329,694923,959611,374501,425996,387202,386965,895264,640655,293693,297949,55565,605886,425074,696281,142621,309118,142185,877986,727355,742910,839660,484653,55828,423592,130733,527751,767489,453985,49497,585199,651897,109417,849602,152089,327984,543504,245456,942597,483713,313810,692068,477215,549731,318905,73810,459803,915743,757623,728140,684987,15596,831972,488677,331475,16059,130624,542524,975244,649582,974326,290568,299236,997613,475880,548113,85735,867703,499767,674001,665596,725475,325163,126649,187371,514842,363975,126699,236622,155278,313854,448028,859345,477212,509380,183198,211217,978015,681805,14583,363457,578577,114404,886605,670317,134046,80368,235036,636813,965046,121028,136476,949557,410458,216676,276856,421340,534913,389306,97644,84780,609074,72405,750727,850908,143487,364714,874084,418607,206227,86850,724059,941159,267116,997646,145917,306502,764626,499327,811099,252030,480052,183345,250051,287006,880092,566549,29809,839361,970718,311345,933045,283455,361738,295204,938728,412948,444659,22535,140489,191697,437972,389243,847112,139419,194383,933443,712602,920062,138722,188699,772992,609674,114093,935794,363988,205366,164192,196309,288567,1749,138941,914963,515680,356596,794168,609680,47753,983259,670565,917059,140631,886513,875037,625372,862161,97095,241820,760993,93226,954168,608282,953877,926674,477422,928775,605527,651686,355050,523257,923432,992322,284578,540584,973270,522477,126034,630061,761443,980766,922269,121853,600041,165428,393378,654525,132791,929492,149544,875563,486938,231974,991720,850959,340909,472402,189510,735042,2973,424771,851245,804394,798575,710729,623892,5825,182226,868243,556013,764873,692092,300730,543043,850444,695357,784254,43734,282010,309203,552625,402465,401112,386646,359759,265190,917148,598312,811206,60845,373921,511613,308212,30139,737918,109200,855566,488897,218124,730465,643452,654214,616722,86264,597488,441619,724351,283558,27997,276463,758794,899919,205349,549649,802234,793570,334696,102978,956851,257051,57949,462368,952967,866800,353882,764930,680700,857127,188533,999594,112596,872139,571475,758845,179795,877493,942814,461521,943040,906313,309798,711985,839649,177931,559074,224113,682828,680469,722624,709174,340156,220966,374817,63988,867606,296774,942133,927676,408482,987698,542760,60319,143507,369395,757244,953903,889440,806602,157261,952013,156165,219062,265670,35826]
		//[71,684,890,415,393,888,25,826,800,555,794,419,878,151,836,713,747,460,368,479,653,886,806,572,290,676,874,843,592,614,490,703,268,998,238,197,395,38,587,143,887,844,301,430,374,38,664,615,93,216,354,907,352,272,322,516,122,171,820,89,891,935,532,49,703,168,58,361,697,164,750,523,696,268,57,326,791,446,869,580,277,483,835,753,790,752,564,792,113,450,398,722,950,580,372,587,59,561,925,34,167,409,605,791,806,309,176,820,475,110,774,276,576,264,418,326,991,455,959,722,848,170,197,189,401,784,547,946,19,93,554,704,67,838,373,734,46,744,904,312,331,851,692,936,925,152,61,334,86,505,62,478,441,958,765,16,150,47,282,724,133,457,912,923,971,182,613,515,558,988,719,229,974,657,907,992,946,840,759,49,45,162,126,789,709,960,992,517,576,605,567,754,319,131,454,4,399,968,43,269,929,195,320,384,767,187,97,363,651,295,520,345,683,126,865,979,141,995,827,395,82,969,92,510,126,345,395,61,793,530,266,643,512,184,782,381,2,729,161,264,134,332,848,536,603,358,839,24,933,773,973,494,799,997,172,665,193,516,714,187,444,470,679,998,182,621,903,964,849,885,673,727,884,212,669,112,579,345,948,933,583,41,774,654,134,138,575,469,750,90,830,395,430,545,275,746,635,911,339,572,931,836,504,193,32,644,113,106,524,527,327,13,34,813,440,849,754,940,770,310,714,854,570,245,465,251,917,436,318,174,598,893,809,874,851,536,786,470,290,947,195,282,506,121,876,110,570,227,232,303,409,187,232,227,299,478,949,307,581,714,405,520,513,662,884,786,933,208,428,99,675,362,983,10,907,310,653,845,345,917,891,498,739,173,167,731,417,115,327,231,38,334,884,230,942,773,580,613,838,709,110,154,702,556,444,272,840,901,298,595,326,859,693,990,346,459,759,107,674,443,127,713,672,31,717,174,580,166,634,614,511,435,922,406,791,307,127,935,63,179,7,344,650,156,190,406,29,914,929,730,492,773,779,668,38,753,297,92,118,270,585,149,479,569,394,302,641]
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, bestTeamScore, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-211/problems/best-team-with-no-conflicts/