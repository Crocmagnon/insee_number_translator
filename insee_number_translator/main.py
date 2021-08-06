import calendar
import datetime
import pprint
import sys

from insee_number_translator.data.cities import CITIES
from insee_number_translator.data.countries import CONTINENTS, COUNTRIES
from insee_number_translator.data.departments import DEPARTMENTS


class InseeData:
    def __init__(self, insee_number: str):
        self.insee_number = insee_number
        if len(insee_number) != 15:
            raise ValueError("insee number must be 15 characters")
        self._gender = insee_number[0]
        self._year = insee_number[1:3]
        self.month = int(insee_number[3:5])
        department = insee_number[5:7]
        city = insee_number[7:10]
        if department in ["97", "98"]:
            department = insee_number[5:8]
            city = insee_number[8:10]

        self.foreign = False
        if department in ["91", "92", "93", "94", "95", "96", "99"]:
            self.foreign = True
            self._country = city
        else:
            self._department = department
            self._city = city

        self.order_of_birth = int(insee_number[10:13])
        self.control_key = int(insee_number[13:])

    @property
    def is_valid(self):
        code = 97 - (
            int(self.insee_number.replace("2A", "19").replace("2B", "18")[:-2]) % 97
        )

        return code == self.control_key

    @property
    def year(self):
        year = int("20" + self._year)
        date = datetime.date(year, self.month, 1)
        if date < datetime.date.today():
            return year
        return year - 100

    @property
    def gender(self):
        if self._gender == "1":
            return "Male"
        return "Female"

    @property
    def gender_short(self):
        if self._gender == "1":
            return "M"
        return "F"

    @property
    def city(self):
        if self.foreign:
            return "unknown"
        return self._city

    @property
    def country(self):
        if self.foreign:
            return self._country
        return "FR"

    @property
    def department(self):
        if self.foreign:
            return "unknown"
        return self._department

    def __str__(self):
        message = [self.insee_number]

        if self.is_valid:
            message.append("The number is valid.")
        else:
            message.append("The number is invalid.")

        message.append(
            f"You're a {self.gender}, born in {calendar.month_name[self.month]}, probably in {self.year}."
        )

        if self.foreign:
            word = "country"
            line = "You're born outside France, "
            country = COUNTRIES.get("99" + self.country, [])
            continent = CONTINENTS.get(self.country[0])
            if len(country) > 1:
                line += f"probably in one of these countries/territories: {', '.join(country)} ({continent})"
            elif len(country) == 1:
                line += f"in {country[0]} ({continent})"
            else:
                line += f"in an unknown country numbered {self.country}"
            message.append(line)
        else:
            word = "city"
            city = CITIES.get(self.department + self.city)
            department_nice = DEPARTMENTS[self.department]
            if city:
                message.append(
                    f"You're born in {city['name']}, {city['zip_code']} ({department_nice}), France."
                )
            else:
                message.append(
                    f"You're born in a unknown city numbered {self.city}, which is located in {department_nice}, in France."
                )

        message.append(
            f"You're the {self.order_of_birth}th to be born in this {word} on this month."
        )

        return "\n".join(message)

    def to_dict(self):
        data = {
            "insee_number": self.insee_number,
            "is_valid": self.is_valid,
            "gender": self.gender_short,
            "month": self.month,
            "year": self.year,
            "order_of_birth": self.order_of_birth,
        }

        if self.foreign:
            data["foreigner"] = {
                "countries_names": COUNTRIES.get("99" + self.country, []),
                "country_code": self.country,
                "continent": CONTINENTS.get(self.country[0]),
            }
        else:
            city = CITIES.get(self.department + self.city, dict())
            data["french"] = {
                "department_name": DEPARTMENTS[self.department],
                "city_insee_code": self.city,
                "department_code": self.department,
                "city_name": city.get("name"),
                "zip_code": city.get("zip_code"),
            }

        return data


def main():
    numbers = ["269059913116714", "168127982980507"]
    if len(sys.argv) > 1:
        numbers = sys.argv[1:]
    for number in numbers:
        data = InseeData(number)
        print(data)
        pprint.pprint(data.to_dict())
        print("\n\n")


if __name__ == "__main__":
    main()
