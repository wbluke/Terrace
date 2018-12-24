# Testfile

### This is cpp code

```cpp
#include<iostream>
#include<cstring>
using namespace std;

int student, soulmateCount;
bool soulmate[10][10];
bool matching[10];

int countPairing(bool matching[10]) {

	int first = -1;
	for (int i = 0; i < student; i++) {
		if (!matching[i]) {
			first = i;
			break;
		}
	}

	//base case
	if (first == -1) return 1;

	int ret = 0;
	for (int soul = first + 1; soul < student; soul++) {
		if (!matching[soul] && soulmate[first][soul]) {
			matching[first] = matching[soul] = true;
			ret += countPairing(matching);
			matching[first] = matching[soul] = false;
		}
	}
	return ret;
}

int main(void) {

	int num, f1, f2;
	cin >> num;

	for (int i = 0; i < num; i++) {

		cin >> student >> soulmateCount;

		memset(soulmate, false, sizeof(soulmate));
		memset(matching, false, sizeof(matching));

		for (int j = 0; j < soulmateCount; j++) {
			cin >> f1 >> f2;
			soulmate[f1][f2] = soulmate[f2][f1] = true;
		}
		cout << countPairing(matching) << endl;

	}


	return 0;
}
```

ok. please.
